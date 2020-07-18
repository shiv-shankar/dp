package dejson

import (
    "log"
    "encoding/json"
    "errors"
    apdm "spca/apd/models"
)

// DeJson : Contains json instance info
type DeJson struct{}


// NewDeJsonHdl : Provides Json handle upon registration with the interface
func NewDeJsonHdl() *DeJson {
    log.Println("NewDeJsonHdl: Json Handle .....")
    DeJsonHdl := DeJson{}
    return &DeJsonHdl
}


// EncodeUsrMsg : Encoding of the User Message structure
func (de *DeJson)EncodeUsrMsg(ip apdm.MsgHdr) ([]byte, error){

//    log.Printf("EncodeUsrMsg input: %+v\n", ip)
    
    msg, err := json.Marshal(ip)

    return msg, err
}

// DecodeUsrMsg : Decoding of the User Message
func (de *DeJson)DecodeUsrMsg(msg []byte) (apdm.MsgHdr, error){

//    log.Printf("DecodeUsrMsg  Msg: %+v\n", msg)

    var cmd apdm.MsgHdr

    err := json.Unmarshal(msg, &cmd)

    return cmd, err 
}

// EncodeCmdConf : Encoding of the Command Conf message
func (de *DeJson)EncodeCmdConf(ip interface{}) ([]byte, error){

//    log.Printf("EncodeCmdConf input: %+v\n", ip)
    msg, err := json.Marshal(ip.(apdm.CmdConf))
    return msg, err 
}

// DecodeCmdConf : Decoding of the Command Config message
func (de *DeJson)DecodeCmdConf(msg []byte) (interface{}, error){

  //  log.Printf("DecodeCmdConf  Msg: %+v\n", msg)

    var cmd apdm.CmdConf

    err := json.Unmarshal(msg, &cmd)

    return cmd, err 
}

// EncodeCmdStatus : Encoding of the Command Status message
func (de *DeJson)EncodeCmdStatus(ip interface{}) ([]byte, error){

 //   log.Printf("EncodeCmdStatus IP: %+v\n", ip)
    msg, err := json.Marshal(&ip)
    return msg, err
}

// DecodeCmdStatus : Decoding of the Command Status message
func (de *DeJson) DecodeCmdStatus(msg []byte) (interface{}, error){

//    log.Printf("DecodeCmdStatus msg : %+v", msg)
    var cmd apdm.CmdOpStatus

    err := json.Unmarshal(msg, &cmd)

    return cmd, err 
}


// EncodeTest ...
func (td *Td1) EncodeTest(ip interface{}) ([]byte, error){

    switch td.D1 {
        case 9:
            // test data encapsulation
            log.Println("Test data encode required")
            ips := ip.(Td1)

            var jmsg []byte
            jmsg, err := json.Marshal(ips)
            if err != nil {
                log.Println(err)
            }

            return jmsg, nil
        default:
            log.Println("Unsupported Encode")
    }

    return nil, errors.New("Json Encode Failed")
}

type Td1 struct {
    D1 int
    D2 string
}

