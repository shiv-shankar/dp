package mongodb


import (
    "log"
    "time"
    "errors"
    apdm "spca/apd/models"
)



type DbRepo struct {
    dbConn     *redis.Client
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
	
    opts, err := redis.ParseURL(redisURL)
    
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opts)
	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}

    Repo.dbConn = client

    log.Println("Redis DB Repository Creation .....")

    return &Repo, nil
}

func (r *DbRepo) FindFirstAvailableCmdForProcess(status string) (*apdm.CmdIp, error) {

/*
    cmdinfo := apdm.CmdIp{
        Id : c.id,
        Cmd : c.cmd,
        CmdType : c.cmdtype,
        DepCmd  : c.depcmd,
        DepCmdParam: c.depcmdparam,
    }
*/
	key := "commands"
	data, err := r.client.HGetAll(key).Result()
	if err != nil {
		return nil, errors.Wrap(err, "Commands.handler.Find")
	}
	if len(data) == 0 {
		return nil, errors.New("Entry not found")
	}

    // identify the first entry with status submitted

    return &cmdinfo, nil
}


func (r *DbRepo) GetCmdStatus() (apdm.CmdOpStatus, error){
    
    op := apdm.CmdOpStatus{
        CmdSubmitted:0,
        CmdRunning:0,
        CmdCompleted:0,
        CmdError:0,
    }


    return op, nil
}


func (r *DbRepo) SetCmdStatusProcessing(id uint) error {

    var err error

    err = nil

    return err
}


func (r *DbRepo) SetCmdStatusRunning(id uint) error {

    var err error

    err = nil

    log.Printf("SetCmdStatusRunning : %+v\n", id)


    return err
}


func (r *DbRepo) SetCmdStatusComplete(id uint) error {

    var err error

    err = nil


    return err
}


func (r *DbRepo) SetCmdStatusError(id uint) error {

    var err error

    err = nil


    return err
}


func (r *DbRepo) StoreCmdInfo(cip apdm.CmdConf)  (uint,error) {

    
    log.Printf("StoreCmdInfo: Input %+v\n", cip)
	key := "commands"
	data := map[string]interface{}{
            "CmdType": cip.CmdType,
            "Cmd":     cip.Cmd,
            "CmdParam": cip.CmdParam,
            "DepCmd": cip.DepCmd,
            "DepCmdParam":cip.DepCmdParam,
	}
	_, err := r.client.HMSet(key, data).Result()
	if err != nil {
		return errors.Wrap(err, "repository.Redirect.Store")
	}
    id := 0
    log.Println("StoreCmdInfo Complete")
    return id, nil
}

func (r *DbRepo) GetCmdInfo(id uint)  (apdm.CmdIp, error) {

    log.Println("Get Command Info FOr Id:", id)
/*
    cmdinfo := apdm.CmdIp{
        Id : c.id,
        Cmd : c.cmd,
        CmdParam: c.cmdparam,
        CmdType : c.cmdtype,
        DepCmd  : c.depcmd,
        DepCmdParam  : c.depcmdparam,
    }
*/
    return cmdinfo, nil
}






