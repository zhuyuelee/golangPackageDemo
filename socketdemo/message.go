package socketdemo

import (
	"GoSql/socketdemo/utils"
	"errors"
	"fmt"
)

//HeadLength header length
const HeadLength uint32 = 0xC

//Message Message
type Message struct {
	Msg string
	ID  uint32
	CID CommandID
}

//CommandID CommandID
type CommandID uint32

const (
	//CONNECT 链接
	CONNECT CommandID = iota + 1
	//CONTENT 内容发送
	CONTENT
	//ERROR 错误
	ERROR
	//CLOSE 关闭
	CLOSE
)

//GetMessageSlice GetMessageSlice
func GetMessageSlice(data []byte) []Message {
	i := uint32(0)
	l := uint32(len(data))
	msgs := make([]Message, 0)

	for {
		head, err := NewHeader(data[i : i+HeadLength])
		if err != nil {
			fmt.Println("read head error,err=", err)
			break
		}
		body := string(data[i+HeadLength : i+head.Lenth])
		msgs = append(msgs, NewMessage(body, head.ID, head.CID))
		i += head.Lenth
		if i >= l {
			break
		}
	}

	return msgs
}

//NewMessage NewMessage
func NewMessage(msg string, id uint32, cid CommandID) Message {
	return Message{
		Msg: msg,
		ID:  id,
		CID: cid,
	}
}

//SetMessage SetMessage
func SetMessage(data []byte, header Header) Message {
	return Message{
		Msg: string(data),
		ID:  header.ID,
		CID: header.CID,
	}
}

//GetMessage GetMessage
func (m *Message) GetMessage() []byte {
	b := []byte(m.Msg)
	h := getHeader(len(b), m.ID, m.CID)
	return append(h, b...)
}

//getHeader getHeader  cid 1 链接 2 发送消息 3关闭链接
func getHeader(length int, id uint32, cid CommandID) (head []byte) {
	head = make([]byte, 0)
	head = append(head, utils.IntToBytes(uint32(length)+HeadLength)...)
	head = append(head, utils.IntToBytes(id)...)
	head = append(head, utils.IntToBytes(uint32(cid))...)
	return
}

//Header Header
type Header struct {
	Lenth uint32
	// totalLength uint
	ID  uint32
	CID CommandID
}

//NewHeader NewHeader
func NewHeader(head []byte) (header Header, err error) {
	header = Header{}
	lenth := uint32(len(head))
	if lenth != HeadLength {
		return header, errors.New("header length error")
	}
	header.Lenth = utils.BytesToInt(head[:4])
	if header.Lenth == 0 {
		return header, fmt.Errorf("data len is 0 head=%v", head)
	}
	header.ID = utils.BytesToInt(head[4:8])
	header.CID = CommandID(utils.BytesToInt(head[8:]))

	return
}

// //GetTotalLenth GetLenth
// func (h Header) GetTotalLenth() int32 {
// 	if len(h.Head) != 12 {
// 		return 0
// 	}
// 	return utils.BytesToInt(h.Head[:4])
// }

// //GetID GetID
// func (h Header) GetID() int32 {
// 	if len(h.Head) != 12 {
// 		return 0
// 	}
// 	return utils.BytesToInt(h.Head[4:8])
// }

// //GetSID GetSID
// func (h Header) GetSID() int32 {
// 	if len(h.Head) != 12 {
// 		return 0
// 	}
// 	return utils.BytesToInt(h.Head[8:])
// }

// //Body Body
// type Body struct {
// 	Body []byte
// }

// //NewBody NewBody
// func NewBody(body string) Body {
// 	return Body{
// 		Body: []byte(body),
// 	}
// }

// //GetBody GetBody
// func (b *Body) GetBody() string {
// 	return string(b.Body)
// }
