package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main()  {
	//Reader()
	ReadByte()
}

func ReadByte()  {
	var ch []byte
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please input your name age number: ")
	err := inputReader.UnreadByte()
	if err != nil{
		fmt.Println(err)
		return
	}
	inputReader.Read(ch)

	fmt.Println(ch)
}




func Reader(){
	FOREND:
		for  {
			readMenu()
			var ch string
			fmt.Scanln(&ch)
			var (
				data []byte
				err error
			)

			switch strings.ToLower(ch) {
			case "1":
				fmt.Println("请输入不多于9个字符，以回车结束：")
				data, err = ReadFrom(os.Stdin, 11)
			case "2":
				file,err:= os.Open("./01txt")
				if err != nil{
					fmt.Println("打开文件 01txt 错误:", err)
					continue
				}
				data, err = ReadFrom(file, 11)
				file.Close()
			case "3":
				data, err = ReadFrom(strings.NewReader("from string"), 12)
			case "4":
				data, err = ReadFrom(bytes.NewReader([]byte("from string")), 12)
			case "5":
				resp ,err := http.Get("http://www.baidu.com")
				if err != nil{
					fmt.Println("打开文件 01txt 错误:", err)
					continue
				}
				data, err = ReadFrom(resp.Body, 1024)
				resp.Body.Close()
			case "b":
				fmt.Println("返回上级菜单")
				break FOREND
			case "q":
				fmt.Println("程序退出")
				os.Exit(0)
			default:
				fmt.Println("输入错误！")
				continue
			}
			if err != nil {
				fmt.Println("数据读取失败，可以试试从其他输入源读取！")
			} else {
				fmt.Printf("读取到的数据是：%s\n", data)
			}
		}
}





func ReadFrom(reader io.Reader, num int)([]byte, error)  {
	buf := make([]byte, num)
	n,err := reader.Read(buf)
	if n>0{
		return buf[:n],nil
	}
	return buf,err
}
func readMenu()  {
	menu := `""
"*******从不同来源读取数据*********
*******请选择数据源，请输入：*********
	1 表示 标准输入
	2 表示 普通文件
	3 表示 从字符串
	4 表示 从字节
	5 表示 从网络
	b 返回上级菜单
	q 退出
***********************************`
	fmt.Println(menu)
}


