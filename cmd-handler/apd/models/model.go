package apdmodel

type CmdIp struct {
    Id          uint
    CmdType     uint
    Cmd         string
    CmdParam    string
    DepCmd      string
    DepCmdParam string
}

type CmdOpStatus struct {
    CmdSubmitted    uint
    CmdProcessing   uint
    CmdRunning      uint
    CmdCompleted    uint
    CmdError        uint
}

type CmdConf struct {
    Service     string
    CmdType     uint
    Cmd         string
    CmdParam    string
    DepCmd      string
    DepCmdParam string

}

type MsgHdr struct {
   MsgType uint
   MsgBody  []byte
}


type CmdResult struct {
    Id uint
    Err error
}

const (
    SERVICEA = "SERVICE-1"
    CMDA    = "CMDA"
    CMDB    = "CMDB"
    CMDC    = "CMDC"
    CMDPARAM1 = "CMD-PARAMA" 
    CMDPARAM2 = "CMD-PARAMB"
)

const (
    INDCMD = 0
    DEPCMD = 1
    MSGCMDCONF = 2
    MSGCMDSTATUS = 3
)

