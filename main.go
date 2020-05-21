package main

import (
	"exercise/cache"
	"fmt"
	"time"
)

func main() {
	c, err := cache.NewLFUCache(3)

	if err != nil{
		fmt.Println(err)
		return
	}
	go c.Put(1, "4444")
	go c.Put(1, "1111")
	go c.Put(1, "2222")
	go c.Put(1, "66666")
	go c.Put(1, "3333222")
	c.Print()
	fmt.Println("---------------------------")

	//c.Put(2, "25544")
	//c.Print()
	//fmt.Println("---------------------------")
	//c.Get(2)
	//c.Put(3, "77344")
	//c.Print()
	//fmt.Println("---------------------------")
	//c.Put(5, "4554656")
	//c.Print()
	//c.Get(3)
	//fmt.Println("---------------------------")
	//c.Put(5, "046888")
	//c.Print()
	//fmt.Println("---------------------------")
	//c.Put(7, "6633")
	//c.Print()
	//fmt.Println("---------------------------")
	//c.Put(8, "9454554")
	//c.Print()
	//fmt.Println("---------------------------")
	////
	////

	//fmt.Println("---------------------------")
	time.Sleep(time.Minute*20)
}
