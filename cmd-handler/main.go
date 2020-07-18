package main

import (
    "log"
    "fmt"
    "io/ioutil"
    "time"
    "os"
    "encoding/json"
    cgw "spca/apd/cmdorchestrator"
    apdm "spca/apd/models"
)

type Cmds struct {
    Cmds []apdm.CmdConf `json:"cmds,omitempty"`
}

var cmds Cmds

func main() {

    log.Println("Command Handler: Started")

    ic := make(chan []byte, 10)
    oc := make(chan []byte, 10)
    tc := make(chan int,2)

    go cgw.CmdHandler(ic, oc)
    go CmdMenu(ic, tc)
    // get the command list from the command file
    jsonFile, err := os.Open("cmd.json")
    byteValue, _ := ioutil.ReadAll(jsonFile)

    err = json.Unmarshal(byteValue, &cmds)
    if err != nil {
        fmt.Println("error:", err)
        os.Exit(1)
    }

    for i :=0; i < len(cmds.Cmds); i++ {

        fmt.Printf("COmand %+v\n", cmds.Cmds[i])

    }
    
    ticker := time.NewTicker(10 * time.Second)
Exit:
    for {

        select{
        case  msg := <- oc:

            var msgst apdm.MsgHdr

            json.Unmarshal(msg,&msgst)

            log.Printf("Msg From CMD handler for status: %+v\n", msgst)

            if msgst.MsgType == apdm.MSGCMDSTATUS{
                var cmdstatus apdm.CmdOpStatus


                json.Unmarshal(msgst.MsgBody, &cmdstatus)


                log.Printf("CMD Status From CMD handler for status: %+v\n", &cmdstatus)
                log.Println(">>>>>>>>>> Command Status <<<<<<<<<<<<")


                log.Printf("Submitted: %d \n",cmdstatus.CmdSubmitted)
                log.Printf("Processing: %d \n",cmdstatus.CmdProcessing)
                log.Printf("Running: %d \n",cmdstatus.CmdRunning)
                log.Printf("Completed: %d ",cmdstatus.CmdCompleted)
                log.Printf("Error:%d\n",cmdstatus.CmdError)
                log.Println(">>>>>>>>>>>>>>>>>> Command Status >>>>>>>>>>>")
            }
        case <-tc:
            break Exit

        case <-ticker.C:

        }

    }

    close(ic)
    close(oc)
    close(tc)

    log.Println("Applicaiton: Stopped")

}

func CmdMenu(ic chan []byte, tc chan int) {

    n:=0
    c:=0

    ELOOP:    
    for {
        fmt.Println("CMD Menu")
        fmt.Println("1 - Send N Conf Commands")
        fmt.Println("2 - Send CMD Status Commands")
        fmt.Println("3 - exit Application")

        fmt.Scanf("%d",&c)

        switch c {
        case 1:
            fmt.Println("Enter number of conf commands to send")
            fmt.Scanf("%d",&n)
            SendConfCommands(ic, n)
        case 2:
            SendCmdStatusMsg(ic)
        case 3:
            fmt.Println("Terminating the Application")
            tc <- 1
            break ELOOP
        default:
            fmt.Println("Unsupported option")
        }
    }
}

func SendCmdStatusMsg(ip chan[]byte){

        // construct msg header
        msgst := apdm.MsgHdr{
               MsgType: apdm.MSGCMDSTATUS,
        }

        msg, _ := json.Marshal(msgst)

        ip <- msg
}


func SendConfCommands(ip chan[]byte, nc int){

    c := 0

    for i := 0 ; i < nc; i++ {


        cnfcmd, _ := json.Marshal(cmds.Cmds[c])


        // construct msg header
        msgst := apdm.MsgHdr{
               MsgType: apdm.MSGCMDCONF,
               MsgBody: cnfcmd,
        }

        msg, _ := json.Marshal(msgst)

        ip <- msg

        c += 1

        if c >= 7 {
            c = 0
        }

    }

}


