package dbgw


import (
    "fmt"
    "testing"
    apd "spca/apd/models"
)


func  TestDbGwOps(t *testing.T){
    
    fmt.Println("DB Gataway Test cases")

    DbGwInit("localhost", "mock", 30)

    cip := apd.CmdConf{
        Cmd:"DBGW-CMDA ParamA",
        CmdType: 0,
    }

    GetDbHandle().StoreCmdInfo(cip)
    opst, _ := GetDbHandle().GetCmdStatus() 

    fmt.Printf("Status of the Cmd op : %+v\n", opst)

    cip.Cmd = "DBGW-CMDB ParamA"
    cip.DepCmd = "DBGW-CMDA ParamA"
    cip.CmdType = 1

    GetDbHandle().StoreCmdInfo(cip)
    
    opst, _ = GetDbHandle().GetCmdStatus() 
    fmt.Printf("Status of the Cmd op : %+v\n", opst)

    op, _ := GetDbHandle().FindFirstAvailableCmdForProcess("CMD-Submitted")
    fmt.Printf("FIrst submitted available %+v\n",op)

    GetDbHandle().SetCmdStatusRunning(op.Id)

    opst, _ = GetDbHandle().GetCmdStatus() 
    fmt.Printf("Status of the Cmd op : %+v\n", opst)


}


