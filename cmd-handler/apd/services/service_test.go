package service


import (
//    "fmt"
    "testing"
    "log"
//    cgw "spca/iif/cmdgw/if"
)


func TestServiceOps(t *testing.T){

    log.Println("Service Tests")

    s := NewServiceCmds()

    c := s.NewCmd(CMDA, "CMDA-Params")

    c.Execute()
    
    c = s.NewCmd(CMDB, "CMDB-Params")

    c.Execute()

    c = s.NewCmd(CMDC, "CMDC-Params")

    c.Execute()


}


