/**
双向链表
实现向尾部添加；向头部添加； 删除一个entry；从尾部弹出；从头部弹出；根据key查找；打印

 */
package double_linked_list

import (
	"errors"
	"fmt"
	"sync"
)

//链表节点的结构
type Entry struct {
	sync.RWMutex
	key interface{}
	value interface{}
	prev *Entry
	next *Entry
}

func NewEntry(key, value interface{}) *Entry {
	return &Entry{key:key, value:value}
}

//返回这个节点的key
func (e *Entry)Key() interface{} {
	return e.key
}
//返回这个节点的value
func (e *Entry)Value() interface{} {
	e.RLock()
	defer e.RUnlock()
	return e.value
}

func (e *Entry)SetValue(value interface{})  {
	e.Lock()
	defer e.Unlock()
	e.value = value
}

//DoubleLinkedList
type DoubleLinkedList struct {
	sync.RWMutex
	head *Entry   //头节点地址
	tail *Entry	  // 尾节点地址
	len uint16 // 长度
	cap uint16    // 容量
}
//从返回长度
func (d *DoubleLinkedList)Len() uint16 {
	d.RLock()
	defer d.RUnlock()
	return d.len
}
//从返回容量
func (d *DoubleLinkedList)Cap() uint16 {
	return d.cap
}

//初始化一个 双向链表
func NewDoubleLinkedList(cap uint16) *DoubleLinkedList {
	if cap == 0 { //如果 cap为0 ，则
		cap = 1<<16 - 1  // 1111 1111 1111 1111
	}
	return &DoubleLinkedList{cap:cap}
}
//判断链表是否已满
func (d *DoubleLinkedList)isFull() error {
	if  d.cap  == d.len{
		return errors.New("该链表已经满")
	}
	return nil
}
//做长度的运算
func  (d *DoubleLinkedList)setLen(isAdd bool)  {
	//加锁
	d.Lock()
	defer d.Unlock()
	if isAdd {
		d.len++
	}else {
		d.len--
	}
}
//向尾部添加
func (d *DoubleLinkedList)Append(entry *Entry)  error{

	if err := d.isFull(); err !=nil{
		return err
	}
	if d.len == 0 {
		d.head = entry
	}else {
		entry.prev = d.tail
		d.tail.next = entry
	}
	d.tail = entry
	d.setLen(true)
	return nil
}
//向头部添加
func (d *DoubleLinkedList)Head(entry *Entry) error {
	if err := d.isFull(); err !=nil{
		return err
	}

	if d.len == 0 {
		d.tail = entry
	}else {
		entry.next = d.head
		d.head.prev = entry
	}
	d.head = entry
	d.setLen(true)
	return nil
}
//从尾部弹出
func (d *DoubleLinkedList)Pop() *Entry {
	entry := d.tail
	if entry == nil {
		return entry
	}

	if entry.prev == nil{
		d.head, d.tail = nil, nil
	}else {
		entry.prev.next = nil
		d.tail = entry.prev
	}
	d.setLen(false)
	return entry
}
//从头部弹出
func (d *DoubleLinkedList)Shift() *Entry {
	entry := d.head
	if entry == nil {
		return entry
	}

	if entry.next == nil{
		d.head, d.tail = nil, nil
	}else {
		entry.next.prev = nil
		d.head = entry.next
	}
	d.setLen(false)
	return entry
}

//删除节点
func (d *DoubleLinkedList)Delete(entry *Entry) error {
	switch  {
	case entry == nil:
	case entry == d.head:
		d.Shift()
	case entry == d.tail:
		d.Pop()
	default:
		entry.prev.next = entry.next
		entry.next.prev = entry.prev
		d.setLen(false)
	}
	return nil
}

//查找
func (d *DoubleLinkedList)Find(key interface{}) []*Entry {
	entry := d.head
	var ret []*Entry
	for entry != nil {
		if key == entry.key{
			ret = append(ret, entry)
		}
		entry = entry.next
	}
	return ret
}
//查找一个
func (d *DoubleLinkedList)FindOne(key interface{}) *Entry {
	d.Lock()
	entry := d.head
	for entry != nil {
		if key == entry.key{
			break
		}
		entry = entry.next
	}
	d.Unlock()
	return entry
}
//循环打印链表
func (d *DoubleLinkedList)Print()  {
	entry := d.head
	for entry != nil {
		fmt.Printf("key 值为 %v,value 值为 %v \n", entry.key, entry.value)
		entry = entry.next
	}
}



