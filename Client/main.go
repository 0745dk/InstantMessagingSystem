package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func ReceiveMessages(co net.Conn) {
	for {
		var message [128]byte
		n, err := co.Read(message[:])
		if err != nil {
			fmt.Println("Receive message error:", err)
			return
		}
		fmt.Printf("\n%s", string(message[:n]))
	}
}
func main() {
	conn, err := net.Dial("tcp4", fmt.Sprintf("%s:%d", "127.0.0.1", 8888))
	if err != nil {
		fmt.Println("Connect Server Failed :", err)
		return
	}
	defer conn.Close()
	//监听服务器发送回来的信息
	go ReceiveMessages(conn)

	//本地就向服务器输入数据。
	for {
		fmt.Print(">")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("input error:", err)
		}
		input = strings.TrimSuffix(input, "\n")
		conn.Write([]byte(input))
	}
}
