package main

import (
	"fmt"
	"net"
	"os"
)

const IpAddr  = "127.0.0.1:8080"

var isQuit chan bool
//发送消息
func Send(conn net.Conn)  {
	input := make([]byte, 1024)
	for  {
		n, err := os.Stdin.Read(input)
		if err != nil{
			fmt.Println("os.Stdin.Read err", err)
			return
		}
		if n == 0 {
			continue
		}
		conn.Write(input[:n])

		if string(input[:n]) == "exit"{
			break
		}
	}
	isQuit<-true   //退出程序
}
//读取消息
func Read(conn net.Conn)  {
	output := make([]byte, 1024)
	for  {
		n, err := conn.Read(output)
		if err != nil{
			fmt.Println("conn.Read", err)
			return
		}
		fmt.Printf(string(output[:n]))
	}
}
// 连接 server ，并循环处理消息
func main() {
	conn,err := net.Dial("tcp", IpAddr)

	if err != nil{
		fmt.Println("net dial err", err)
		return
	}
	defer conn.Close()

	go Send(conn)
	go Read(conn)

	<-isQuit
}


