package cmdgw

import (
    "log"
    "testing"
    "time"
    apdm "spca/apd/models"
)


func TestCmdGwOps(t *testing.T){

    log.Println("Command Gw Tests")

    var cip apdm.CmdIp
    opc := make(chan apdm.CmdResult, 2)

    CmdGwInit()

    cip.Id = 3
    cip.CmdType = apdm.INDCMD
    cip.Cmd = "CMDA"
    cip.CmdParam = "Param-1"

    CmdScheduler(&cip, opc)
    time.Sleep(1 * time.Second)

    log.Println("Command Gw Tests Complete ", <- opc)

}

