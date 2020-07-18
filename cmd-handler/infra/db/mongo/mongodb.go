package mongodb


import (
    "log"
    "time"
    "errors"
    apdm "spca/apd/models"
)



type DbRepo struct {
    dbConn    *mongo.Client
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

    ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(DbURL))
	if err != nil {
		return nil, err
	}
	
    err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	
	Repo.dbConn = client

    log.Println("Mock DB Repository Creation .....")

    return &Repo, nil
}

func (r *DbRepo) FindFirstAvailableCmdForProcess(status string) (*apdm.CmdIp, error) {

    ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

    var cmdinfo = apdm.CmdIp


	collection := r.client.Database(r.Dbname).Collection("commands")
	filter := bson.M{"status": status}
	err := collection.FindOne(ctx, filter).Decode(&cmdinfo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("repository.Redirect.Find")
		}
		return nil, errors.Wrap(err, "repository.CommandsFind")
	}


    return &cmdinfo, nil
}


func (r *DbRepo) GetCmdStatus() (apdm.CmdOpStatus, error){
    

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

    ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	
    collection := r.client.Database(r.Dbname).Collection("commands")
	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"CmdType": cip.CmdType,
			"Cmd":     cip.Cmd,
			"CmdParam": cip.CmdParam,
            "DepCmd": cip.DepCmd,
            "DepCmdParam":cip.DepCmdParam,
		},
	)
	if err != nil {
		return errors.Wrap(err, "repository.Redirect.Store")
	}    
    
    log.Println("StoreCmdInfo Complete")
    return id, nil
}

func (r *DbRepo) GetCmdInfo(id uint)  (apdm.CmdIp, error) {

    log.Println("Get Command Info For Id:", id)
    return cmdinfo, nil
}






