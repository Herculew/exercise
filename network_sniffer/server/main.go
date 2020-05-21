/**
网络嗅探工具


本想通过语言把网卡设置为混杂模式，最终没研究出来，选择了直接shell操作了
type Ifreq struct {
	Name [unix.IFNAMSIZ]byte
	Data uintptr
}
var ifreq Ifreq
_,_,errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd),
	uintptr(syscall.SIOCSIFFLAGS), uintptr(unsafe.Pointer(&ifreq)))

if errno != 0{
fmt.Println(errno.Error())
	return
}
*/
package main

import (
	"exercise/network_sniffer/util"
	"fmt"
	"syscall"
)


func Htons(i uint16) uint16 {
	return (i<<8)&0xff00 | i>>8
}

func main() {
	util.Promisc("promisc")   //开启网卡混杂模式
	//int(Htons(0x800))
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_RAW)
	util.CheckError(err)

	for {
		//go func(f *os.File) {
		//123456
		//}(f)
		buf := make([]byte, 1500)
		n,err:=syscall.Read(fd, buf)


		fmt.Println(n)
		fmt.Println(err)
		//header := network.ParseHeader(buf[0:14])
		//
		//if header.DestinationAddress.String() != "00:00:00:00:00:00"{
		//	fmt.Println(header)
		//}

	}

}
