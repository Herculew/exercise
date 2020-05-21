package util

import (
	"fmt"
	"os"
	"os/exec"
)

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s \n", err.Error())
		os.Exit(1)
	}
}

//开启网卡混杂模式
// promisc 开启
//-promisc 取消
func Promisc(s string)  {
	cmd := exec.Command("sudo","ifconfig", "wlp5s0", s)
	err := cmd.Run()
	CheckError(err)
}