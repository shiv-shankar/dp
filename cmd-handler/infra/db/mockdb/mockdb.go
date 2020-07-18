package mockdb


import (
    "log"
    "time"
    "errors"
    apdm "spca/apd/models"
)


type CmdInfo struct {
    id          uint
    cmdtype     uint
    cmd         string
    cmdparam    string
    depcmd      string
    depcmdparam string
    status      string
    submitTime  int64
    comptime    int64
    exelog      string
    result      string
}



type mkdb struct {
    cmdInfoTbl  map[uint]CmdInfo
    lastProcEntry uint
    freeId        uint
    stats   apdm.CmdOpStatus
}




type DbRepo struct {
    MkdbConn    mkdb
    Timeout     int
    Dbname      string
    dbUrl       string
}

// DB handle
var DbHandle *DbRepo

// NewMockRepository Provides new DB handle
func NewMockRepository(DbURL string, Db string, DbTimeout int) (*DbRepo, error) {


    //Mock db does not need connections but initialization
    Repo := DbRepo{
        Timeout : DbTimeout,
        Dbname: Db,
        dbUrl: DbURL,
    }

    // Initialize the data structure
    Repo.MkdbConn.cmdInfoTbl = make(map[uint]CmdInfo)
    Repo.MkdbConn.lastProcEntry = 0
    Repo.MkdbConn.freeId = 0

    log.Println("Mock DB Repository Creation .....")

    return &Repo, nil
}

func (r *DbRepo) FindFirstAvailableCmdForProcess(status string) (*apdm.CmdIp, error) {

    if r.MkdbConn.lastProcEntry >= r.MkdbConn.freeId {
//        return nil,errors.New("No Entries available") 
        return nil,nil
    }
    c := r.MkdbConn.cmdInfoTbl[r.MkdbConn.lastProcEntry]

    cmdinfo := apdm.CmdIp{
        Id : c.id,
        Cmd : c.cmd,
        CmdType : c.cmdtype,
        DepCmd  : c.depcmd,
        DepCmdParam: c.depcmdparam,
    }

    c.status =  "CMD-Processing"
    r.MkdbConn.cmdInfoTbl[r.MkdbConn.lastProcEntry]  = c
    r.MkdbConn.stats.CmdProcessing += 1
    r.MkdbConn.stats.CmdSubmitted -= 1


    r.MkdbConn.lastProcEntry += 1

    return &cmdinfo, nil
}


func (r *DbRepo) GetCmdStatus() (apdm.CmdOpStatus, error){
    
    op := apdm.CmdOpStatus{
        CmdSubmitted:0,
        CmdRunning:0,
        CmdCompleted:0,
        CmdError:0,
    }

    op.CmdSubmitted = r.MkdbConn.stats.CmdSubmitted
    op.CmdProcessing = r.MkdbConn.stats.CmdProcessing
    op.CmdRunning = r.MkdbConn.stats.CmdRunning
    op.CmdCompleted = r.MkdbConn.stats.CmdCompleted
    op.CmdError = r.MkdbConn.stats.CmdError

    return op, nil
}


func (r *DbRepo) SetCmdStatusProcessing(id uint) error {

    var err error

    err = nil

    c := r.MkdbConn.cmdInfoTbl[id]

    if c.status == "CMD-Submitted" {
        c.status = "CMD-Processing"
        r.MkdbConn.cmdInfoTbl[id]  = c
        r.MkdbConn.stats.CmdProcessing += 1
        r.MkdbConn.stats.CmdSubmitted -= 1
        r.MkdbConn.lastProcEntry += 1

    } else {
        err = errors.New("Incorrect status update: Failed Op")
    }

    return err
}


func (r *DbRepo) SetCmdStatusRunning(id uint) error {

    var err error

    err = nil

    log.Printf("SetCmdStatusRunning : %+v\n", id)

    c := r.MkdbConn.cmdInfoTbl[id]
    log.Printf("SetCmdStatusRunning : %+v c :%+v \n", id,c)

    if c.status == "CMD-Processing" {
        c.status = "CMD-Running"
        r.MkdbConn.cmdInfoTbl[id]  = c
        r.MkdbConn.stats.CmdProcessing -= 1
        r.MkdbConn.stats.CmdRunning += 1
    } else {
        err = errors.New("Incorrect status update: Failed Op")
    }

    return err
}


func (r *DbRepo) SetCmdStatusComplete(id uint) error {

    var err error

    err = nil

    c := r.MkdbConn.cmdInfoTbl[id]

    log.Printf("SetCmdStatusComplete : %+v c :%+v \n", id,c)
    if c.status == "CMD-Running" {
        c.status = "CMD-Completed"
        c.comptime = time.Now().UTC().Unix()
        r.MkdbConn.cmdInfoTbl[id]  = c
        r.MkdbConn.stats.CmdCompleted += 1
        r.MkdbConn.stats.CmdRunning -= 1
    } else {
        err = errors.New("Incorrect status update: Failed Op")
    }

    return err
}


func (r *DbRepo) SetCmdStatusError(id uint) error {

    var err error

    err = nil

    c := r.MkdbConn.cmdInfoTbl[id]

    log.Printf("SetCmdStatusError : %+v c :%+v \n", id,c)
    if c.status == "CMD-Running" {
        c.status = "CMD-Error"
        c.comptime = time.Now().UTC().Unix()
        r.MkdbConn.cmdInfoTbl[id]  = c
        r.MkdbConn.stats.CmdError += 1
        r.MkdbConn.stats.CmdRunning -= 1
    } else {
        err = errors.New("Incorrect status update: Failed Op")
    }

    return err
}


func (r *DbRepo) StoreCmdInfo(cip apdm.CmdConf)  (uint,error) {

    
    log.Printf("StoreCmdInfo: Input %+v\n", cip)
    c:= CmdInfo{
        id:r.MkdbConn.freeId,
        cmd: cip.Cmd,
        cmdparam: cip.CmdParam,
        cmdtype: cip.CmdType,
        depcmd: cip.DepCmd,
        depcmdparam: cip.DepCmdParam,
        status: "CMD-Submitted",
        submitTime:time.Now().UTC().Unix(),
    }

    r.MkdbConn.cmdInfoTbl[r.MkdbConn.freeId] = c

    id:= r.MkdbConn.freeId
    log.Printf("inserted element i: %d val : %+v\n", r.MkdbConn.freeId,r.MkdbConn.cmdInfoTbl[r.MkdbConn.freeId])
    r.MkdbConn.freeId += 1
    r.MkdbConn.stats.CmdSubmitted += 1
    
    log.Println("StoreCmdInfo Complete")
    return id, nil
}

func (r *DbRepo) GetCmdInfo(id uint)  (apdm.CmdIp, error) {

    log.Println("Get Command Info FOr Id:", id)
    c := r.MkdbConn.cmdInfoTbl[id]

    cmdinfo := apdm.CmdIp{
        Id : c.id,
        Cmd : c.cmd,
        CmdParam: c.cmdparam,
        CmdType : c.cmdtype,
        DepCmd  : c.depcmd,
        DepCmdParam  : c.depcmdparam,
    }

    return cmdinfo, nil
}






