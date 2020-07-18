package dexml

import (
        "fmt"
        "log"
        "testing"
        "errors"
        apdm "spca/apd/models"
)


func TestDeXmlSetup(t *testing.T){

    fmt.Println("Test 1: Data Encode Setting to Xml Setup")
    xmlHdl := NewDeXmlHdl()

    if xmlHdl == nil {
        fmt.Println("Error in getting the Xml handle")
        t.Fatal(errors.New("Get Xml Handle Failed"))
    }

    fmt.Println("Got Xml Handle")

}



func TestDeXmlOper(t *testing.T){
    fmt.Println("Test 2: Xml operations Setup/Encode/Decode")


    xmlHdl := NewDeXmlHdl()

    if xmlHdl == nil {
        fmt.Println("Error in getting the Xml handle")
        t.Fatal(errors.New("Get Xml Handle Failed"))
    }

    fmt.Println("Xml Handle Get Success:")

    fmt.Println("Xml Encode Test:")
    td := Td1{9,"Test-CMD"}
    t1 := Td1{D1:9, D2:"SGN"}
    msg, err := td.EncodeTest(t1)

    fmt.Printf("Xml Encoded Msg: %+v Error: %+v\n",msg, err)

    fmt.Println("Xml Encode Test Success:")
    fmt.Println("Xml Decode Test:")
/*
    xmlHdl.EncodeCmdConf("Encode-Cmd-Conf")
    xmlHdl.DecodeCmdConf([]byte(`{"Name":"Decode-Cmd-Conf"}`))
    xmlHdl.EncodeCmdStatus("Encode-Cmd-Status")
    xmlHdl.DecodeCmdStatus([]byte(`{"Name":"Decode-Cmd-Status"}`))

*/
    cmd := apdm.CmdConf{
        Service:"SERVICE-1",
        Cmd:"CMDA",
        CmdType:apdm.INDCMD,
        CmdParam:"Param-1",
    }

    msg, _ = xmlHdl.EncodeCmdConf(cmd)
    log.Println(string(msg))
    
    cmd1, _ := xmlHdl.DecodeCmdConf(msg)
    log.Println(cmd1.(apdm.CmdConf))

    cmdst := apdm.CmdOpStatus{
        CmdSubmitted: 3,
        CmdProcessing: 2,
        CmdRunning: 4,
        CmdCompleted: 9,
        CmdError:1,
    }

    msg,_ = xmlHdl.EncodeCmdStatus(cmdst)
    log.Println(string(msg))
    
    cmdst1, _ := xmlHdl.DecodeCmdStatus(msg)
    log.Println(cmdst1.(apdm.CmdOpStatus))


    fmt.Println("Xml Decode Test Success:")

}

