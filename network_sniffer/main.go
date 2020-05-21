package main

import "fmt"

// Htons ...
func Htons(i uint16) uint16 {
	return (i<<8)&0xff00 | i>>8
}

func main()  {
	fmt.Println(int(Htons(0x800)))
}