package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

/**
利用tcp socket编程，上传文件 ----服务端
1.监听服务，开启端口，

 */
const IpAddr  = "127.0.0.1:8080"
func main() {
	listener, err := net.Listen("tcp", IpAddr)

	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}

	defer listener.Close()

	for  {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err: ", err)
			continue
		}

		go HandleUpload(conn)
	}

}
//
func HandleUpload(conn net.Conn)  {
	defer conn.Close()

	bufioR := bufio.NewReader(conn)
	buf := make([]byte, 1024)

	n, err := bufioR.Read(buf)

	if err != nil {
		fmt.Println("bufioR.Read err: ", err)
		return
	}
	filename := string(buf[:n])
	bufioW := bufio.NewWriter(conn)
	_, err1 := conn.Write([]byte("ok"))

	if err1 != nil {
		fmt.Println("onn.Write ok err: ", err1)
		return
	}
	//创建文件
	fmt.Println("recv filename: ", filename)

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("bufioR.Read err: ", err)
		return
	}
	for  {
		n, err := conn.Read(buf)
		if err != nil && err != io.EOF{
			fmt.Println("bufioR.Read content err: ", err)
			return
		}
		if n == 0 {
			fmt.Println("n == 0 文件接收完毕")
			bufioW.WriteString("file compelte")
			return
		}
		file.Write(buf[:n])
	}

}
