package main

import (
	"fmt"
	"net"
	"golang.org/x/net/icmp"
)

func main() {
	netaddr, err := net.ResolveIPAddr("ip4", "127.0.0.1")
	if err !=nil{
		fmt.Println(err.Error())
		return
	}
	IPConn, err := net.ListenIP("ip4:icmp", netaddr)
	if err !=nil{
		fmt.Println(err.Error())
		return
	}
	buf := make([]byte, 1024)
	for {
		n, addr, _ := IPConn.ReadFrom(buf)
		msg,_:= icmp.ParseMessage(1,buf[0:n])
		fmt.Println(n, addr, msg.Type,msg.Code,msg.Checksum)
	}
}
