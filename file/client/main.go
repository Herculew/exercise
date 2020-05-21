package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

/*
利用tcp socket编程，上传文件 ---- 客户端

1.读取命令行里面的参数，进行参数验证
2.获取上传文件的信息，并验证上传文件是否存在
3.向服务端传送文件名，并获取回复，为下一步上传文件做准备
4.循环读取文件，并传输到服务端。
5.测试，解决bug
是否考虑 连接复用，或者连接池
*/

const IpAddr  = "127.0.0.1:8080"

var (
	arg = os.Args
	exit = make(chan bool)
)

func main() {
	if l:=len(arg); l != 2 {
		fmt.Println("usage is err ")
		return
	}
	fileInfo, err := os.Stat(arg[1])
	if err != nil {
		fmt.Println("os.Stat err: ", err)
		return
	}

	go uploadFile(fileInfo, exit)

	<-exit
}
//开启一个协程上传文件
func uploadFile(fileInfo os.FileInfo, exit chan<- bool)  {

	total, filename, size := 0, fileInfo.Name(), fileInfo.Size()
	fmt.Println("upload filename is: ", filename)
	//连接服务端
	conn, err := net.Dial("tcp",IpAddr)

	//不管错误还是传输成功完成，都要结束客户端
	defer func() {exit <- true}()

	if err != nil {
		fmt.Println("net.Dial err: ", err)
		return
	}
	defer conn.Close()

	//上传文件名给服务端
	_, err1 := conn.Write([]byte(filename))
	if err1 != nil {
		fmt.Println("conn.Write filename err: ", err1)
		return
	}

	//收取服务端的回复
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read `ok` err: ", err)
		return
	}

	if "ok" != string(buf[:n]){
		fmt.Println("conn.Read not equal `ok` err: ")
		return
	}

	file ,err := os.Open(arg[1])

	if err != nil {
		fmt.Println("os.Open err: ", err)
		return
	}

	for  {
		n, err := file.Read(buf)
		if err != nil &&  err != io.EOF{
			fmt.Println("file.Read err: ", err)
			return
		}
		if n == 0{
			fmt.Println("file.Read complete: ",filename )
			return
		}
		total += n
		fmt.Printf("file.upload  %f: \n", float64(total)/float64(size)*100.00)
		conn.Write(buf[:n])
	}
}