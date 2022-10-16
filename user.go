package main

import (
	"fmt"
	. "net"
)

type User struct {
	Uid     int
	Name    string
	Addr    string
	Connect Conn
	Channel chan string //用来接受handler协程发送过来的消息，handler是处理client送给服务器消息的协程
}

func NewUser(connect Conn, id int) *User {
	user := &User{
		Uid:     id,
		Name:    connect.RemoteAddr().String(), //RemoteAddr()返回一个连接中，远程方的地址
		Addr:    connect.RemoteAddr().String(),
		Connect: connect,
		Channel: make(chan string),
	}
	//new完之后接着开协程发消息
	go user.InformClientGoroutine()
	return user
}

// InformClientGoroutine 负责给客户端发消息，将客户收到的消息回显给客户端
func (t User) InformClientGoroutine() {
	var message string
	for {
		message = <-t.Channel
		_, err := t.Connect.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Server->User:", t.Uid, " error:cannot write messages to client.")
			fmt.Println("message:", message)
		}
	}
}
