package dexml

import (
    "log"
    "encoding/xml"
    "errors"
    apdm "spca/apd/models"
)

// DeXml : Contains xml instance info
type DeXml struct{}


// NewDeXmlHdl : Provides Xml handle upon registration with the interface
func NewDeXmlHdl() *DeXml {
    log.Println("NewDeXmlHdl: Xml Handle .....")
    DeXmlHdl := DeXml{}
    return &DeXmlHdl
}


// EncodeUsrMsg : Encoding of the User Message structure
func (de *DeXml)EncodeUsrMsg(ip apdm.MsgHdr) ([]byte, error){

    log.Printf("EncodeCmdConf input: %+v\n", ip)
    
    msg, err := xml.Marshal(ip)

    return msg, err
}

// DecodeUsrMsg : Decoding of the User Message
func (de *DeXml)DecodeUsrMsg(msg []byte) (apdm.MsgHdr, error){

    log.Printf("DecodeCmdConf  Msg: %+v\n", msg)

    var cmd apdm.MsgHdr

    err := xml.Unmarshal(msg, &cmd)

    return cmd, err 
}

// EncodeCmdConf : Encoding of the Command Conf message
func (de *DeXml)EncodeCmdConf(ip interface{}) ([]byte, error){

    log.Printf("EncodeCmdConf input: %+v\n", ip)
    msg, err := xml.Marshal(ip.(apdm.CmdConf))
    return msg, err 
}

// DecodeCmdConf : Decoding of the Command Config message
func (de *DeXml)DecodeCmdConf(msg []byte) (interface{}, error){

    log.Printf("DecodeCmdConf  Msg: %+v\n", msg)

    var cmd apdm.CmdConf

    err := xml.Unmarshal(msg, &cmd)

    return cmd, err 
}

// EncodeCmdStatus : Encoding of the Command Status message
func (de *DeXml)EncodeCmdStatus(ip interface{}) ([]byte, error){

    log.Printf("EncodeCmdStatus IP: %+v\n", ip)
    msg, err := xml.Marshal(&ip)
    return msg, err
}

// DecodeCmdStatus : Decoding of the Command Status message
func (de *DeXml) DecodeCmdStatus(msg []byte) (interface{}, error){

    log.Printf("DecodeCmdStatus msg : %+v", msg)
    var cmd apdm.CmdOpStatus

    err := xml.Unmarshal(msg, &cmd)

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
            jmsg, err := xml.Marshal(ips)
            if err != nil {
                log.Println(err)
            }

            return jmsg, nil
        default:
            log.Println("Unsupported Encode")
    }

    return nil, errors.New("Xml Encode Failed")
}

type Td1 struct {
    D1 int
    D2 string
}

