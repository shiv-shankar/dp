package service


import (
        "log"
        cgw "spca/iif/cmdgw/if"
)

// Service 
type Service  struct {
    Name string
    CmdCount int
}

// NewCmdService
func NewServiceCmds(s string) *Service {
/*    
    clist := make([]string, 3)

    clist = append(clist, CMDA)
    clist = append(clist, CMDB)
    clist = append(clist, CMDC)
*/
    return &Service{
        Name:s,
        CmdCount:0,
    }
}

func (s *Service)NewCmd(ct string, param string) cgw.CmdGw {

    switch ct {
    case  CMDA:
        return s.NewCmdA(param)
    case  CMDB:
        return s.NewCmdB(param)
    case  CMDC:
        return s.NewCmdC(param)
    default:
        log.Println("Unsupported Command")
        return nil
    }
}

//NewCmdA Setups the command with the parameters provided by the user 
func (s *Service) NewCmdA(p string) cgw.CmdGw {
    return &CmdA{
        Params:p,
    }
}

// NewCmdB Constructs the Command with the parameters 
func (s *Service) NewCmdB(p string) cgw.CmdGw {
    return &CmdB{
        Params:p,
    }
}

// NewCmdC Constructs the Command with the parameters 
func (s *Service) NewCmdC(p string) cgw.CmdGw {
    return &CmdC{
        Params:p,
    }
}


