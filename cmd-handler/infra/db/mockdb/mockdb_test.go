package mockdb

import (
    "fmt"
    "testing"
//    dbgw "spca/iif/dbgw"
    apdm "spca/apd/models"
)


func TestMockDbOp(t *testing.T) {
    
    fmt.Println("Mock Db Test Cases")

    mdbhdl, _ := NewMockRepository("localhost", "MOCKDB", 20)

    fmt.Println("DB repo ", mdbhdl)

    cip := apdm.CmdConf{
        Service:"SERVCE-1",
        Cmd: "CMDA",
        CmdType: apdm.INDCMD,
        CmdParam:"Param-1",

    }

    mdbhdl.StoreCmdInfo(cip)
    
    opst, _ := mdbhdl.GetCmdStatus()
    
    fmt.Printf("Status of the Cmd op : %+v\n", opst)

    cip.Cmd = "CMDB"


    mdbhdl.StoreCmdInfo(cip)
    
    opst, _ = mdbhdl.GetCmdStatus()
    
    fmt.Printf("Status of the Cmd op : %+v\n", opst)

    for i:=0 ; i < 10; i++ {
    
    //    cip.Cmd = fmt.Sprintf("CMDB-%d %s",  i,  " Param")
        mdbhdl.StoreCmdInfo(cip)
    
        opst, _ = mdbhdl.GetCmdStatus()
    
        fmt.Printf("Status of the Cmd op : %+v\n", opst)
    }

    opst, _ = mdbhdl.GetCmdStatus()
    
    fmt.Printf("Status of the Cmd op : %+v\n", opst)

    op, _ := mdbhdl.FindFirstAvailableCmdForProcess("CMD-Submitted")
    fmt.Printf("FIrst submitted available %+v\n",op)

    opst, _ = mdbhdl.GetCmdStatus()
    fmt.Printf("Status of the Cmd op : %+v\n", opst)
    

    for i := 0; i < 8 ; i++ { 

        op, _ := mdbhdl.FindFirstAvailableCmdForProcess("CMD-Submitted")
        fmt.Printf("FIrst submitted available %+v\n",op)

        opst, _ = mdbhdl.GetCmdStatus()
        fmt.Printf("Status of the Cmd op : %+v\n", opst)

        mdbhdl.SetCmdStatusRunning(op.Id)
        
        opst, _ = mdbhdl.GetCmdStatus()
        fmt.Printf("Status of the Cmd op : %+v\n", opst)
        
        
        mdbhdl.SetCmdStatusComplete(op.Id)
        
        opst, _ = mdbhdl.GetCmdStatus()
        fmt.Printf("Status of the Cmd op : %+v\n", opst)
    }
        
    op, _ = mdbhdl.FindFirstAvailableCmdForProcess("CMD-Submitted")
    fmt.Printf("FIrst submitted available %+v\n",op)
    
    opst, _ = mdbhdl.GetCmdStatus()
    fmt.Printf("Status of the Cmd op : %+v\n", opst)
        
        mdbhdl.SetCmdStatusRunning(op.Id)
        
    opst, _ = mdbhdl.GetCmdStatus()
    fmt.Printf("Status of the Cmd op : %+v\n", opst)
    
    mdbhdl.SetCmdStatusError(op.Id)
        
    opst, _ = mdbhdl.GetCmdStatus()
    fmt.Printf("Status of the Cmd op : %+v\n", opst)

}



