/*
LRU可以使用 二进制移位, 链表,堆栈 三种方式都可以实现
这里采取了链表的方式
 */
package cache

import (
	"errors"
	"exercise/cache/double_linked_list"
)

type LRUCache struct {
	list *double_linked_list.DoubleLinkedList
}

func NewLRUCache(cap uint16) (*LRUCache, error) {
	if cap == 0{
		return nil, errors.New("容量必须大于0")
	}
	return &LRUCache{list: double_linked_list.NewDoubleLinkedList(cap)}, nil
}

//返回长度
func (lru *LRUCache) Len() uint16 {
	return lru.list.Len()
}
//返回容量
func (lru *LRUCache) Cap() uint16 {
	return lru.list.Cap()
}
//得到key的value， 如果不存在就返回nil
func (lru *LRUCache) Get(key interface{}) *double_linked_list.Entry {
	entry := lru.list.FindOne(key)
	if entry != nil{
		lru.list.Delete(entry)  //删除旧的
		lru.list.Head(entry)
	}
	return entry
}
//放置一个
func (lru *LRUCache) Put(key, value interface{})  {
	//判断是否已经在缓存里面
	if entry := lru.list.FindOne(key);entry !=nil{
		lru.list.Delete(entry)  //删除旧的
	} else if lru.list.Len() == lru.list.Cap(){
		lru.list.Pop()   //删除尾部节点
	}

	lru.list.Head(double_linked_list.NewEntry(key, value)) //头部追加进去

	return
}
//打印一下
func (lru *LRUCache)Print()  {
	lru.list.Print()
}