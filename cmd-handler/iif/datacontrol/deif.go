package dceif

import (
    "log"
    dcejs "spca/infra/de/json"
    dcexml "spca/infra/de/xml"
    apdm "spca/apd/models"
)


const (
    DTJSON = "json"
    DTXML = "xml"
)



var DceifHdl map[string]DataCtrlIf

// DataCtrlInit Initialize the interface
func DataCtrlInit(){
    log.Println("Data Controller Initialization ")
    DceifHdl = make(map[string]DataCtrlIf)
    log.Println("Data Controller Initialization ..... Complete")
}


/* DataCtrlI: ALl the interfaces that application uses to perform 
 encoding and decoding. Each structure packing and/or unpacking will 
 need to have interface definition
*/
type DataCtrlIf interface{
    EncodeUsrMsg(ip apdm.MsgHdr) ([]byte, error)
    DecodeUsrMsg(msg []byte) (apdm.MsgHdr, error)
    EncodeCmdConf(ip interface{}) ([]byte, error)
    DecodeCmdConf(msg []byte) (interface{}, error)
    EncodeCmdStatus(ip interface{}) ([]byte, error)
    DecodeCmdStatus(msg []byte) (interface{}, error)
}

// DataCtrlSetup : Registration of the Data transport 
func DataCtrlSetup(dceType string){
    log.Println("Data Control : Setup")
    
    t := dceType

    switch t {
    case "json":
        DceifHdl[dceType] = dcejs.NewDeJsonHdl()
        log.Println("JSON Registration Complete")

    case "xml":
        DceifHdl[dceType] = dcexml.NewDeXmlHdl()
        log.Println("XML Registration Complete")
    default:
        log.Println("UnSupported Data Transport")
    }
    log.Println("Data Control : Setup ..... Complete")

}

// GetHandle : Provides the handle to the Data Controller interface
func GetHandle(dceType string) (DataCtrlIf) {
    return DceifHdl[dceType]
}


