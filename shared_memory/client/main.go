package main

type ShmEntry struct {
	canRead chan bool
	buf [1024]byte
}
// todo 后面完成
//https://blog.nlogn.cn/posts/share-memory-golang/
//https://studygolang.com/articles/10203
//https://www.iminho.me/wiki/docs/understand_linux_process/golang-memory.md

func main()  {
//	ShmEntry := &ShmEntry{

}
