package dejson

import (
        "fmt"
        "log"
        "testing"
        "errors"
        apdm "spca/apd/models"
)


func TestDeJsonSetup(t *testing.T){

    fmt.Println("Test 1: Data Encode Setting to Json Setup")
    jsonHdl := NewDeJsonHdl()

    if jsonHdl == nil {
        fmt.Println("Error in getting the Json handle")
        t.Fatal(errors.New("Get Json Handle Failed"))
    }

    fmt.Println("Got Json Handle")

}



func TestDeJsonOper(t *testing.T){
    fmt.Println("Test 2: Json operations Setup/Encode/Decode")


    jsonHdl := NewDeJsonHdl()

    if jsonHdl == nil {
        fmt.Println("Error in getting the Json handle")
        t.Fatal(errors.New("Get Json Handle Failed"))
    }

    fmt.Println("Json Handle Get Success:")

    fmt.Println("Json Encode Test:")
    td := Td1{9,"Test-CMD"}
    t1 := Td1{D1:9, D2:"SGN"}
    msg, err := td.EncodeTest(t1)

    fmt.Printf("Json Encoded Msg: %+v Error: %+v\n",msg, err)

    fmt.Println("Json Encode Test Success:")
    fmt.Println("Json Decode Test:")
/*
    jsonHdl.EncodeCmdConf("Encode-Cmd-Conf")
    jsonHdl.DecodeCmdConf([]byte(`{"Name":"Decode-Cmd-Conf"}`))
    jsonHdl.EncodeCmdStatus("Encode-Cmd-Status")
    jsonHdl.DecodeCmdStatus([]byte(`{"Name":"Decode-Cmd-Status"}`))

*/
    cmd := apdm.CmdConf{
        Service:"SERVICE-1",
        Cmd:"CMDA",
        CmdType:apdm.INDCMD,
        CmdParam:"Param-1",
    }

    msg, _ = jsonHdl.EncodeCmdConf(cmd)
    log.Println(string(msg))
    
    cmd1, _ := jsonHdl.DecodeCmdConf(msg)
    log.Println(cmd1.(apdm.CmdConf))

    cmdst := apdm.CmdOpStatus{
        CmdSubmitted: 3,
        CmdProcessing: 2,
        CmdRunning: 4,
        CmdCompleted: 9,
        CmdError:1,
    }

    msg,_ = jsonHdl.EncodeCmdStatus(cmdst)
    log.Println(string(msg))
    
    cmdst1, _ := jsonHdl.DecodeCmdStatus(msg)
    log.Println(cmdst1.(apdm.CmdOpStatus))


    fmt.Println("Json Decode Test Success:")

}

