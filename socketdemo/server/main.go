package main

import (
	"GoSql/socketdemo"
	"fmt"
	"net"
)

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("net.Listen start error ,err=", err)
	}
	defer server.Close()
	fmt.Println("server starting")
	//发送消息
	go send(msgChan)
	go close()

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("server.Accept error ,err=", err)
		}
		go connect(conn)
	}
}

var clients map[uint32]net.Conn
var closeclients chan uint32
var msgChan chan socketdemo.Message

func init() {
	clients = make(map[uint32]net.Conn)
	msgChan = make(chan socketdemo.Message, 10)
	closeclients = make(chan uint32, 10)
}

// 链接
func connect(conn net.Conn) {
	// defer conn.Close()
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("client conn error, err=", err)
			conn.Close()
		}
	}()

	var data []byte = make([]byte, 12)
	n, err := conn.Read(data)
	if err != nil {
		fmt.Println("read head error,err=", err)
		conn.Close()
		return
	}
	if n == 0 {
		conn.Close()
		return
	}
	fmt.Println("connect data", data)
	head, err := socketdemo.NewHeader(data[:socketdemo.HeadLenth])
	if err != nil {
		fmt.Println("read head error,err=", err)
		conn.Close()
		return
	}
	fmt.Println("head data", head)
	if head.CID == socketdemo.CONNECT {
		clients[head.ID] = conn
		msgChan <- socketdemo.NewMessage("", head.ID, socketdemo.CONNECT)
		go client(conn, head.ID)
		return
	}
	conn.Close()
}

//关闭消息
func close() {
	for {
		id := <-closeclients
		if conn, ok := clients[id]; ok {
			delete(clients, id)
			conn.Close()
		}

	}
}

//client 接收消息
func client(conn net.Conn, id uint32) {
	defer conn.Close()
	defer func() {
		delete(clients, id)
		err := recover()
		if err != nil {
			fmt.Println("client conn error, err=", err)
		}
	}()
	name := conn.RemoteAddr().String()
	fmt.Printf("%s connecting\n", name)
	for {
		var data []byte = make([]byte, 2048)
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println("read head error,err=", err)
			break
		}
		if n > 0 {
			go read(name, data[:n])
		}
		if n == 0 {
			break
		}
	}
}

func read(name string, data []byte) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error=", err)
		}
	}()

	msgs := socketdemo.GetMessageSlice(data)
	closeID := uint32(0)
	for index, msg := range msgs {
		msgChan <- msg
		fmt.Printf("id=%d msg=%s index=%d \n", msg.ID, msg.Msg, index)
		if msg.CID == socketdemo.CLOSE {
			closeID = msg.ID
		}
	}
	if closeID > 0 {
		closeclients <- closeID
	}

}

//发送消息
func send(msgChan <-chan socketdemo.Message) {
	for {
		msg := <-msgChan
		if conn, ok := clients[msg.ID]; ok {
			//发送消息
			conn.Write(msg.GetMessage())
		}

	}
}
