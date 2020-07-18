package cmdhandler

import (
    "log"
    "time"
    cgw "spca/iif/cmdgw"
    dcegw "spca/iif/datacontrol"
    dbgw  "spca/iif/dbgw"
    apdm "spca/apd/models"
)

type CmdHandle  struct {
    ipChan  <-chan []byte
    opChan  chan<- []byte
    CmdOpChan chan apdm.CmdResult
    dce     string
    dbUrl   string
    db      string
    dbtimeout int
    maxwload   int
    currload   int
}

var CmdHdl   CmdHandle

// CmdHandler independent routine that will wait for the command over the channel in Json format
func CmdHandler(ipcc <-chan []byte, opcc chan<- []byte) {

    log.Println("Command Handler: Started")
    CmdHdl.ipChan = ipcc
    CmdHdl.opChan = opcc
    CmdHdl.CmdOpChan = make(chan apdm.CmdResult, 10)

    ticker := time.NewTicker(10 * time.Second)


    // Initialize all the interfaces
    CmdHdl.CmdHandlerInit()

    for {
        select {
            case s := <-ipcc:
//            log.Println(s)
                CmdHdl.ProcessCmd(s)
//            opcc <- "CMD Handler: No Error"
            case cmdOp := <-CmdHdl.CmdOpChan:

                //log.Printf("Command output RX %+v\n",cmdOp)

                if cmdOp.Err != nil {

                    dbgw.GetDbHandle().SetCmdStatusError(cmdOp.Id)

                } else {
                    dbgw.GetDbHandle().SetCmdStatusComplete(cmdOp.Id)
                }


                CmdHdl.currload -= 1


                //trigger Pending Commands
                CmdHdl.CmdProcessPendingCommands()


            case <-ticker.C:
//                log.Println("Ticker Fired")
                CmdHdl.CmdHealthMonitor()

        }

    }

}


func (c *CmdHandle) ProcessCmd(usrmsg []byte) error {
    var err error
    err = nil

//    var msg apdm.MsgHdr

    log.Println("Process Cmd: ")

    //Decode the Message
    msg, err := dcegw.GetHandle(c.dce).DecodeUsrMsg(usrmsg)


    switch msg.MsgType {

    case apdm.MSGCMDCONF:    

        //Decode Message Body
        usrCmd, err := dcegw.GetHandle(c.dce).DecodeCmdConf(msg.MsgBody)

        log.Printf("Decode Command: %+v Err %+v\n", usrCmd, err)

        c.ProcessConfCmd(usrCmd.(apdm.CmdConf))

     case apdm.MSGCMDSTATUS:
        c.ProcessStatusCmd()

    }
   
    return err 
}

func (c *CmdHandle) ProcessStatusCmd() error {

    var err error
    err = nil
    log.Println("Process Command Status: ......")
    op, _ := dbgw.GetDbHandle().GetCmdStatus()
    log.Println("OP Status :",op)

    //Encode the Command Status 
    msgbody,err := dcegw.GetHandle(c.dce).EncodeCmdStatus(op)

    msgst := apdm.MsgHdr{
            MsgType:apdm.MSGCMDSTATUS,
            MsgBody:msgbody,
    }

    msg, err := dcegw.GetHandle(c.dce).EncodeUsrMsg(msgst)

    c.opChan <- msg

    return err
}


func (c *CmdHandle) ProcessConfCmd(usrCmd apdm.CmdConf) error {

    var err error
    err = nil

    var id uint
    
    log.Println("ProcessConfCmd: ......")
    
    // Add the command to the data base
    if id,err = dbgw.GetDbHandle().StoreCmdInfo(usrCmd); err != nil {
        log.Println(err)
        return err
    }


    // check for the work load. in case free send the command for execution
    if c.currload >= c.maxwload {
        log.Println("Current workload is MAX !!!!!!!!!!!")
        return err
    }
    
    ci,_ := dbgw.GetDbHandle().GetCmdInfo(id)

    log.Printf("Get result for id: %d Command Info: %+v\n", id, ci)

    dbgw.GetDbHandle().SetCmdStatusProcessing(id)

    // Send the command for command scheduling
    cgw.CmdScheduler(&ci, c.CmdOpChan)

    dbgw.GetDbHandle().SetCmdStatusRunning(id)


    c.currload += 1

   return err 


}



func (c *CmdHandle) CmdHandlerInit() {
    log.Println("Command Handler Initialization")

    c.CmdHandlerGetEnv()
    c.currload = 0

    // Data Repository Initialization
    log.Println("Data Repository - Complete")


    // Adapter and interface initialization
    c.CmdHandlerIIfInit()
    log.Println("Adapter and interface initialization - Complete")

}

func (c *CmdHandle) CmdHandlerGetEnv() {
    log.Println("Environment Information Initialization")
    //c.dce = os.Getenv("TXRX")
    //c.dbUrl = os.Getenv("URL_DB")
    //c.db = os.Getenv("DB") 
    //c.dbtimeout = os.Getenv("DB_TIMEOUT")
    c.dce = "json"
    c.dbUrl = "localhost"
    c.db = "mock"
    c.dbtimeout = 20
    c.maxwload = 3

    log.Printf("Environment variables DBURL: %s DB: %s Db Timerout: %d DCE: %s ", c.dbUrl, c.db, c.dbtimeout, c.dce)
}

func (c *CmdHandle) CmdHandlerIIfInit() {
    log.Println("CmdHandlerIIfInit: Start")
    
    // Command Gateway 
    cgw.CmdGwInit()
    log.Println("Command Gateway Initialization - Complete") 

    // Data Controller init
    dcegw.DataCtrlInit()

    // Data Transport Initialization
    dcegw.DataCtrlSetup(c.dce)
    log.Println("Data Transport Initialization - Complete") 

    // Database gateway 
    dbgw.DbGwInit(c.dbUrl, c.db, c.dbtimeout)
    log.Println("Database GAteway Initialization - Complete") 

   log.Println("CmdHandlerIIfInit: End")
}



// CmdHealthMonitor Health monitoring for the command handler ...
func (c *CmdHandle) CmdHealthMonitor(){
//    log.Println("CmdHealthMonitor: Start")

    //Check if any of the commands have been running for ever avg timout time 20s
//    log.Println("Monitored Cmd Processing Health")

    // Check the system load if less send more commands for processing
    c.CmdProcessPendingCommands()

//    log.Println("CmdHealthMonitor: Complete")
}

func (c *CmdHandle)CmdProcessPendingCommands() {
    //var err error
    var ci *apdm.CmdIp

//    log.Println("CmdProcessPendingCommands: Check for worker load")
    
    if c.currload < c.maxwload {

//        log.Println("Processing not at full capacity")
        for i := 0; i < (c.maxwload - c.currload); i++ {

            // Get the submitted command
            if ci , _ = dbgw.GetDbHandle().FindFirstAvailableCmdForProcess("CMD-Submitted"); ci == nil {

                break
            }

//            log.Printf("FindFirstsubmit >>>> %+v\n", ci)
            log.Printf("Scheduling Command Id %d:\n", ci.Id)
            //schedule for command execution
            cgw.CmdScheduler(ci, c.CmdOpChan)

            dbgw.GetDbHandle().SetCmdStatusRunning(ci.Id)

            c.currload += 1

        }
    }

}


