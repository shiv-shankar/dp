package cmdgw


import (
    "log"
    apds "spca/apd/services"
    gwif "spca/iif/cmdgw/if"
    apdm  "spca/apd/models"
//    cmdh  "spca/apd/cmdorchestrator"
    
)


// CmdGwHdl ...
var CmdGwHdl  map[string]interface{}

var srvlist []string

// CmdGwInit initializes the datastructure and establishes the connection to the Serives ...
func CmdGwInit(){

    log.Println("Command Gateway Init")
    
    CmdGwHdl = make(map[string]interface{} )
    srvlist = make([]string,1)

    srvlist = append(srvlist, "SERVICE-1")

    //Get the context for each service
    for _, s := range srvlist {
        CmdGwHdl[s] = apds.NewServiceCmds(s)
    }
    log.Println("Command Gateway Init ..... Complete")
}


// CmdScheduler COnstruct the command and then invokes the concurrent command execution ...
func  CmdScheduler(cip *apdm.CmdIp, opc chan<- apdm.CmdResult) error {

    sname := "SERVICE-1"
    log.Printf("CmdScheduler: input : %+v \n", cip)
    s := CmdGwHdl[sname].(*apds.Service)

    // construct the command
    var cmdlist []gwif.CmdGw
//    cmdlist = make([]gwif.CmdGw, 1)

    if cip.CmdType == apdm.DEPCMD {
        depcmd := s.NewCmd(cip.DepCmd,cip.DepCmdParam)
        cmdlist = append(cmdlist, depcmd)
    }
    cmd := s.NewCmd(cip.Cmd, cip.CmdParam)
    cmdlist = append(cmdlist, cmd)

    go CmdGwRunConc(cip, cmdlist, opc)

    log.Println("Trigger Concurrent Exection")
    
    return nil

}

//CmdGwRunConc Concurrent execution of Command and return error propogation
func CmdGwRunConc(cmdinfo *apdm.CmdIp, clist []gwif.CmdGw, opc chan<- apdm.CmdResult) {

//    log.Printf("CmdGwRunConc: Start Command Id: %d \n", cmdinfo.Id)
    var err error
    for _, c := range clist {
        err = c.Execute()
        if err != nil  {
            break
        }
    }

    log.Printf("Command execution returned Err: %+v Command Id: %d", err, cmdinfo.Id)

    cmdResult := apdm.CmdResult{
                    Id :  cmdinfo.Id,
                    Err : err,
                }

    opc <- cmdResult

//    log.Println("CmdGwRunConc: End")
}



