package main

import (
	"GoSql/socketdemo"
	"GoSql/socketdemo/utils"
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("net.Listen start error ,err=", err)
	}
	defer conn.Close()
	fmt.Println("client starting")
	var i = 0

	if !connect(conn) {
		fmt.Println("client connect error")
	} else {
		exit := make(chan bool)
		go read(conn, exit)

		go func() {
			for {
				var input string = fmt.Sprintf("i=%d %s-->", i, utils.RandStr(12, utils.Upper))
				cid := socketdemo.CONTENT
				// if i == 10000 {
				// 	cid = socketdemo.CLOSE
				// }
				msg := socketdemo.NewMessage(input, uint32(1), cid)
				//	fmt.Println("您说：", msg.Msg)
				//发送消息
				conn.Write(msg.GetMessage())
				// if cid == socketdemo.CLOSE {
				// 	break
				// }
				i++
			}
		}()
		<-exit
	}
	fmt.Println("退出了.....")
}

func read(conn net.Conn, exit chan<- bool) {
	defer conn.Close()
	defer func() {
		exit <- true
		err := recover()
		if err != nil {
			fmt.Println("read error,err=", err)
		}
	}()
	reader := bufio.NewReader(conn)
	for {
		head, err := reader.Peek(12)
		if err != nil {
			fmt.Println("reader head error=", err)
			break
		}
		reader.Discard(len(head))

		header, err := socketdemo.NewHeader(head)
		if err != nil {
			fmt.Println("header error=", err)
			break
		}
		if header.CID == socketdemo.CLOSE {
			exit <- true
			break
		} else if header.CID == socketdemo.CONTENT {
			conLen := int(header.Lenth - socketdemo.HeadLength)
			data, err := reader.Peek(conLen)
			if err != nil {
				fmt.Println("data error=", err)
				break
			}
			fmt.Printf("recive id=%d msg=%s\n", header.ID, string(data))
			reader.Discard(conLen)
		}
	}
}

func connect(conn net.Conn) (result bool) {
	msg := socketdemo.NewMessage("", uint32(1), socketdemo.CONNECT)
	connectDate := msg.GetMessage()
	fmt.Println("链接,msg=", msg, "data=", connectDate)
	//发送消息
	conn.Write(msg.GetMessage())

	data := make([]byte, socketdemo.HeadLength)
	n, err := conn.Read(data)
	if err != nil {
		fmt.Println("conn.Read err", err)
		return
	}
	if n == 0 {
		return
	}
	head, err := socketdemo.NewHeader(data[:socketdemo.HeadLength])
	if err != nil {
		fmt.Println("read head error,err=", err)
		conn.Close()
		return
	}
	fmt.Println("connect head", head)
	result = head.CID == socketdemo.CONNECT
	return
}
