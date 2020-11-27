package main

import (
	"GoSql/socketdemo"
	"bufio"
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
	msgChan = make(chan socketdemo.Message, 0)
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
	head, err := socketdemo.NewHeader(data[:socketdemo.HeadLength])
	if err != nil {
		fmt.Println("read head error,err=", err)
		conn.Close()
		return
	}
	fmt.Println("head data", head)
	if head.CID == socketdemo.CONNECT {
		clients[head.ID] = conn
		msgChan <- socketdemo.NewMessage("", head.ID, socketdemo.CONNECT)
		go read(conn, head.ID)
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
			fmt.Printf("%d 关闭了 \n", id)
		}

	}
}

func read(conn net.Conn, id uint32) {
	defer func() {
		closeclients <- id
		err := recover()
		if err != nil {
			fmt.Println("client conn error, err=", err)
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
		if header.CID != socketdemo.CONTENT {
			break
		} else {
			conLen := int(header.Lenth - socketdemo.HeadLength)
			data, err := reader.Peek(conLen)
			if err != nil {
				fmt.Println("data error=", err)
				break
			}
			msg := socketdemo.SetMessage(data, header)
			fmt.Printf("recive id=%d msg=%s\n", header.ID, msg.Msg)
			msgChan <- msg
			reader.Discard(conLen)
		}
	}
}

//send发送消息
func send(msgChan <-chan socketdemo.Message) {
	for {
		msg := <-msgChan
		if conn, ok := clients[msg.ID]; ok {
			//发送消息
			conn.Write(msg.GetMessage())
		}

	}
}
