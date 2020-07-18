package dbgw


import (
        "log"
        apd "spca/apd/models"
        mkdb "spca/infra/db/mockdb"
)

// DbGw Defines all the interface to the database. All interactions happen with these interfaces
type DbGw interface {
    StoreCmdInfo(cip apd.CmdConf) (uint,error)

    FindFirstAvailableCmdForProcess(status string) (*apd.CmdIp, error)

    GetCmdStatus() (apd.CmdOpStatus, error)

    SetCmdStatusProcessing(id uint) error
    SetCmdStatusRunning(id uint) error
    SetCmdStatusComplete(id uint) error
    SetCmdStatusError(id uint) error
    GetCmdInfo(id uint) (apd.CmdIp,error)

}

var DbgwHndl   DbGw

// GetDbHandle : Returns the DB handle 
func GetDbHandle() DbGw {
    return DbgwHndl
}


// DbGWInit Application calls the init funciton with information from environment variables
func DbGwInit(dbUrl string, db string, timeout int) {

    log.Println("Database Gateway init")
    switch db {
        case "mongo":
//            DbgwHndl, _ = NewMongoRepository(dbUrl, db, timeout)
            log.Println("Connecting to Mongo DB .....")
        case "redis":
//            DbgwHndl, _ = NewRedisRepository(dbUrl, db, timeout)
            log.Println("Connecting to Redis DB .....")
        case "mock":
            DbgwHndl, _ = mkdb.NewMockRepository(dbUrl, db, timeout)
            log.Println("Connecting to Mock DB .....")
        default:
            log.Println("Unsupported DB type")
    }
    log.Println("Database Gateway Init ..... Complete")
}





