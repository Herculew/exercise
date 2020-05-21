package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string //用户发送数据的管道
	isQuit chan bool //用户是否退出
	hasData chan bool //用户是否有数据
	Username string      //用户名
	Addr string      //网络地址
}
var (
	//保存在线用户
	online = make(map[string]Client)
	message  = make(chan string)
)

//新开一个协程，转发消息，只要有消息来了，遍历map, 给map每个成员都发送此消息
func Manager() {
	for {
		msg := <-message //没有消息前，这里会阻塞

		//遍历map, 给map每个成员都发送此消息
		for _, cli := range online {
			cli.C <- msg
		}
	}
}
//向客户端发送消息
func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C { //给当前客户端发送信息
		conn.Write([]byte(msg + "\n"))
	}
}
//包装消息体
func MakeMsg(cli Client, msg string)string {
	ret := "[" + cli.Addr + "]" + cli.Username + ": " + msg

	return ret
}
// 从客户端读取消息
func ReadMsgFromClient(cli Client, conn net.Conn, cliAddr string)  {
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf) //todo 也是阻塞
		fmt.Println(111111)
		if err != nil { //对方断开，或者，出问题
			cli.isQuit <- true
			fmt.Println("conn.Read err = ", err)
			return
		}

		msg := string(buf[:n-1]) //如果是在命令下，会多个\n

		switch  {
		case len(msg) == 4 && msg == "list":
			HandleList(cli)
		case len(msg) >= 8 && msg[:6] == "rename":
			HandleRename(cli, msg, cliAddr)
		default:
			message <- MakeMsg(cli, msg)
		}
		cli.hasData <- true //代表有数据
	}
}

//处理改名
func HandleRename(cli Client, msg string, cliAddr string)  {
	// rename|mike
	// todo 如果严谨点，还应该检测语法是否错误
	cli.Username = strings.Split(msg, "|")[1]
	online[cliAddr] = cli
	message <- MakeMsg(cli, "rename ok")

}
//处理是有哪些成员
func HandleList(cli Client)  {
	//遍历map，给当前用户发送所有成员
	userList := "user list:\n"
	for _, tmp := range online {
		userList +=  tmp.Username + "\n"
	}
	message <- MakeMsg(cli, userList)
}
//处理用户链接
func HandleConn(conn net.Conn) {
	defer conn.Close()

	//获取客户端的网络地址
	cliAddr := conn.RemoteAddr().String()

	//创建一个结构体, 默认，用户名和网络地址一样
	cli := Client{make(chan string),make(chan bool),make(chan bool),
		cliAddr, cliAddr}
	//把结构体添加到map
	online[cliAddr] = cli

	//给当前客户端发送信息
	go WriteMsgToClient(cli, conn)
	//广播某用户在线
	message <- MakeMsg(cli, "login")
	//提示，我是谁
	cli.C <- MakeMsg(cli, "I am here")

	//新建一个协程，接收用户发送过来的数据
	go ReadMsgFromClient(cli, conn, cliAddr)

	for {
		//通过select检测channel的流动
		select {
		case <-cli.isQuit:
			delete(online, cliAddr)            //当前用户从map移除
			message <- MakeMsg(cli, "login out") //广播谁下线了
			return
		case <-cli.hasData:
		case <-time.After(60 * time.Second): //设置一个超时，60s后
			delete(online, cliAddr)                     //当前用户从map移除
			message <- MakeMsg(cli, "time out leave out") //广播谁下线了
			return
		}
	}
}

//主函数
func main() {
	//监听
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}

	defer listener.Close()

	//新开一个协程，转发消息，只要有消息来了，遍历map, 给map每个成员都发送此消息
	go Manager()

	//主协程，循环阻塞等待用户链接
	for {
		conn, err := listener.Accept() //todo 阻塞的
		if err != nil {
			fmt.Println("listener.Accept err = ", err)
			continue
		}
		go HandleConn(conn) //处理用户链接
	}

}
