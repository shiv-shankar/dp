package dceif

import (
    "fmt"
    "log"
    "testing"
    apdm "spca/apd/models"
)


func TestDataCtrlr(t *testing.T){

    fmt.Println("Test Data Controller interface")

    DataCtrlInit()

    t1 := "json"
    // registering for JSON Data transport
    DataCtrlSetup(t1)

    t1 = "xml"
    // registering for XML Data transport
    DataCtrlSetup(t1)

    cmd := apdm.CmdConf{
        Service:"SERVICE-1",
        Cmd:"CMDA",
        CmdType:apdm.INDCMD,
        CmdParam:"Param-1",
    }

    msg, _ := GetHandle("json").EncodeCmdConf(cmd)
    log.Println(string(msg))
    
    cmd1, _ := GetHandle("json").DecodeCmdConf(msg)
    log.Println(cmd1.(apdm.CmdConf))

    cmdst := apdm.CmdOpStatus{
        CmdSubmitted: 3,
        CmdProcessing: 2,
        CmdRunning: 4,
        CmdCompleted: 9,
        CmdError:1,
    }

    msg,_ = GetHandle("json").EncodeCmdStatus(cmdst)
    log.Println(string(msg))
    
    cmdst1, _ := GetHandle("json").DecodeCmdStatus(msg)
    log.Println(cmdst1.(apdm.CmdOpStatus))


}


