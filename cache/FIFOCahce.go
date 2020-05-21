package cache

import (
	"errors"
	"exercise/cache/double_linked_list"
)

type FIFOCache struct {
	list *double_linked_list.DoubleLinkedList
}

func NewFIFOCache(cap uint16) (*FIFOCache ,error) {
	if cap == 0{
		return nil, errors.New("容量必须大于0")
	}
	return &FIFOCache{list: double_linked_list.NewDoubleLinkedList(cap)}, nil
}
//返回长度
func (fifo *FIFOCache) Len() uint16 {
	return fifo.list.Len()
}
//返回容量
func (fifo *FIFOCache) Cap() uint16 {
	return fifo.list.Cap()
}
//得到key的value， 如果不存在就返回nil
func (fifo *FIFOCache) Get(key interface{}) *double_linked_list.Entry {
	return fifo.list.FindOne(key)
}
//放置一个
func (fifo *FIFOCache)Put(key,value interface{})  {
	//判断是否已经在缓存里面
	if entry := fifo.list.FindOne(key);entry !=nil{
		fifo.list.Delete(entry)  //删除旧的
	} else if fifo.list.Len() == fifo.list.Cap(){
		fifo.list.Shift()   //删除头节点
	}
	fifo.list.Append(double_linked_list.NewEntry(key, value)) //追加进去
	return
}
//打印一下
func (fifo *FIFOCache)Print()  {
	fifo.list.Print()
}