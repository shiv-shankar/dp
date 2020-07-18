package cmdhandler


import (
    "log"
    "testing"
    apdm "spca/apd/models"
    "encoding/json"
)


func TestCmdHdlerOps(t *testing.T) {

    log.Println("Command Handler Tests: Start")

    ic := make(chan []byte, 10)
    oc := make(chan []byte, 10)

    go CmdHandler(ic, oc)

    cmdst := apdm.CmdConf{
        Service : apdm.SERVICEA,
        CmdType: apdm.DEPCMD,
        Cmd : apdm.CMDA,
        CmdParam : apdm.CMDPARAM1,
        DepCmd : apdm.CMDB,
        DepCmdParam: apdm.CMDPARAM2,
    }

    msgbody, _ := json.Marshal(&cmdst)

    msgst := apdm.MsgHdr{
        MsgType:apdm.MSGCMDCONF,
        MsgBody:msgbody,
    }

    msg, _ := json.Marshal(&msgst)

    log.Printf("Json Marshal : %+v\n", msg)
    ic<- msg


    msgst = apdm.MsgHdr{
        MsgType:apdm.MSGCMDSTATUS,
        //   MsgBody:0,
    }

    msg, _ = json.Marshal(&msgst)

    log.Printf("Json Marshal : %+v\n", msg)

    ic<- msg




    msg = <- oc

    json.Unmarshal(msg,&msgst)

    log.Printf("Msg From CMD handler for status: %+v\n", msgst)
    var cmdstatus apdm.CmdOpStatus

    json.Unmarshal(msgst.MsgBody, &cmdstatus)
    log.Printf("CMD Status From CMD handler for status: %+v\n", &cmdstatus)

    log.Println(">>>>>>>>>> Command Status <<<<<<<<<<<<")
    log.Printf("Submitted: %d Processing: %d Running: %d Completed: %d Error:%d\n",
    cmdstatus.CmdSubmitted,
    cmdstatus.CmdProcessing,
    cmdstatus.CmdRunning,
    cmdstatus.CmdCompleted,
    cmdstatus.CmdError)

    log.Printf("Ouput from the CMD handler %+v\n",msg)
    log.Println(<-oc)

    close(ic)
    close(oc)

    log.Println("Command Handler Tests: End")

}




