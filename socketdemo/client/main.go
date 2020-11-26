package main

import (
	"GoSql/socketdemo"
	"GoSql/socketdemo/utils"
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
		go func() {
			for {
				data := make([]byte, 1024)

				n, err := conn.Read(data)
				if err != nil {
					fmt.Println("conn.Read err", err)
				}
				if n == 0 {
					exit <- true
				}
			}
		}()

		go func() {
			for {
				var input string = fmt.Sprintf("i=%d %s-->", i, utils.RandStr(12, utils.Upper))
				cid := socketdemo.CONTENT
				if i == 10000 {
					cid = socketdemo.CLOSE
				}
				msg := socketdemo.NewMessage(input, uint32(1), cid)
				fmt.Println("您说：", msg.Msg)
				//发送消息
				conn.Write(msg.GetMessage())
				if cid == socketdemo.CLOSE {
					break
				}
				i++
			}
		}()
		<-exit
	}
	fmt.Println("退出了.....")
}

func connect(conn net.Conn) (result bool) {
	msg := socketdemo.NewMessage("", uint32(1), socketdemo.CONNECT)
	connectDate := msg.GetMessage()
	fmt.Println("链接,msg=", msg, "data=", connectDate)
	//发送消息
	conn.Write(msg.GetMessage())

	data := make([]byte, socketdemo.HeadLenth)
	n, err := conn.Read(data)
	if err != nil {
		fmt.Println("conn.Read err", err)
		return
	}
	if n == 0 {
		return
	}
	head, err := socketdemo.NewHeader(data[:socketdemo.HeadLenth])
	if err != nil {
		fmt.Println("read head error,err=", err)
		conn.Close()
		return
	}
	fmt.Println("connect head", head)
	result = head.CID == socketdemo.CONNECT
	return
}
