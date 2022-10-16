package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip               string
	Port             int
	UserMap          map[int]*User
	mapLock          sync.RWMutex
	BoardCastChannel chan string //用于服务器向所有客户端广播消息的管道。一旦有消息传入，服务器就会发送(通过BoardCast方法传入)
}

// NewServer 创建一个新的服务器，返回这个服务器的地址
// 你需要设置它的ip和端口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:               ip,
		Port:             port,
		UserMap:          make(map[int]*User),
		BoardCastChannel: make(chan string),
	}
	return server
}

// HandlerConnection 处理用户连接请求
func (t *Server) HandlerConnection(conn net.Conn, times int) {
	//将用户加入到用户表中
	t.mapLock.Lock()
	t.UserMap[times] = NewUser(conn, times)
	t.mapLock.Unlock()

	fmt.Println("==========================================")
	fmt.Println("uid:", t.UserMap[times].Uid)
	fmt.Println("address:", t.UserMap[times].Addr)
	fmt.Println("connected")
	fmt.Println("==========================================")

	//通知服务器内全体用户该用户上线：
	message := fmt.Sprintf("the user(uid:%d,addr:%s) joined the server, Welcome!", t.UserMap[times].Uid, t.UserMap[times].Addr)
	t.BoardCast(message)
}

func (t *Server) BoardCast(message string) {
	t.BoardCastChannel <- message
}
func (t *Server) BoardCastToOnlineUser() {
	for {
		s := <-t.BoardCastChannel
		t.mapLock.Lock()
		for _, user := range t.UserMap {
			user.Channel <- s
		}
		t.mapLock.Unlock()
	}
}
func (t *Server) Start() {
	//用Server的ip和端口开启监听
	listener, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", t.Ip, t.Port))
	if err != nil {
		fmt.Println("Listen failed:", err)
		return
	} else {
		fmt.Println("listener OK.")
	}
	//关闭监听
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("Close failed:", err)
		}
	}(listener)

	go t.BoardCastToOnlineUser()
	//接受连接
	times := 0
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept Connection failed:", err)
			continue
		}
		//处理连接
		times++
		go t.HandlerConnection(conn, times)
	}
}
