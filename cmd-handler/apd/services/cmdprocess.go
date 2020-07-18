package service


import (
//        "fmt"
        "log"
//        "errors"
        "math/rand"
        "time"
)

// CmdA ...
type CmdA struct {
    Params string
}

//CmdB ...
type CmdB struct {
    Params string
}

// CmdC ...
type CmdC struct {
    Params string
}

const (
    CMDA = "CMDA"
    CMDB = "CMDB"
    CMDC = "CMDC"
)


func (c *CmdA) Execute() error {

    var err error

    rand.Seed(time.Now().UnixNano())
    ri := rand.Intn(10)

    log.Printf("Cmd A Execution Started: Input Paramameters :%v \n", c.Params)
    time.Sleep(time.Duration(ri) * time.Second)
    log.Printf("Cmd A Executed Successfuly: Input Paramameters :%v \n", c.Params)

    return err
}

func (c * CmdB) Execute() error {

    var err error

    rand.Seed(time.Now().UnixNano())
    ri := rand.Intn(10)

    log.Printf("Cmd B Execution Started: Input Paramameters :%v \n", c.Params)
    
    time.Sleep(time.Duration(ri) * time.Second)
    
    log.Printf("Cmd B Executed Successfuly: Input Paramameters :%v \n", c.Params)
    
    return err
}

func (c * CmdC) Execute() error {

    var err error

    rand.Seed(time.Now().UnixNano())
    ri := rand.Intn(10)

    log.Printf("Cmd C Execution Started: Input Paramameters :%v \n", c.Params)
    time.Sleep(time.Duration(ri) * time.Second)
    log.Printf("Cmd C Executed Successfuly: Input Paramameters :%v \n", c.Params)
    return err
}



