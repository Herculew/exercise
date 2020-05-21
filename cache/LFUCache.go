/**
LFU可以使用 二进制移位, 链表都可以实现
这里采取了链表的方式,链表的方式，就必须记录数据访问的频率，但是加了频率，也有一定的问题。
那就是 频率如果一直累加，不做任何其他修改的哈，准确性并不高，比如：
现在有a,b,c,d 四个数据， 前面十分钟，a的访问频率很高，频率数值已经增加到30.
而后面十分钟，并没有在访问a，但是因为a的频率很高，所以一直不会被淘汰，这是不正常的。
，所以每隔60s要清理一次频率。
 */

package cache

import (
	"errors"
	"exercise/cache/double_linked_list"
	"fmt"
	"time"
)

type LFUCache struct {
	freqMap map[interface{}]uint16   // key=>频率值
	list *double_linked_list.DoubleLinkedList
}

//实例化
func NewLFUCache(cap uint16) (*LFUCache, error) {
	if cap == 0{
		return nil, errors.New("容量必须大于0")
	}
	
	LFUCache :=  &LFUCache{
		freqMap: make(map[interface{}]uint16),
		list: double_linked_list.NewDoubleLinkedList(cap),
	}

	go startClear(LFUCache)

	return LFUCache, nil
}
//简单粗暴的解决 每60s清理一次
func startClear(lfuCache *LFUCache) {
	for  {
		select {
		case <-time.After(60 * time.Second):
			if len(lfuCache.freqMap) == 0{
				break
			}
			for key,_ := range lfuCache.freqMap{
				lfuCache.freqMap[key] = 1
			}
			fmt.Printf("60s 后的 %v; \n", lfuCache.freqMap)
		}
	}

}

//得到key的value， 如果不存在就返回nil
func (lfu *LFUCache) Get(key interface{}) *double_linked_list.Entry {

	entry := lfu.list.FindOne(key)
	if entry != nil{
		lfu.freqMap[key]++
	}
	return entry
}
//放置一个
func (lfu *LFUCache) Put(key, value interface{})  {
	//是否在链表中,在链表中，设置值
	if entry := lfu.list.FindOne(key);entry !=nil{
		entry.SetValue(value)
		lfu.freqMap[key]++
		return
	}
	//容量已经满了，需要淘汰一个
	if lfu.list.Len() == lfu.list.Cap(){
		minKey := lfu.getMinFreqKey()  //找到频率最小的key
		lfu.list.Delete(lfu.list.FindOne(minKey))  //删除这个entry
		delete(lfu.freqMap, minKey)
	}

	lfu.list.Append(double_linked_list.NewEntry(key, value)) //把新的节点加入到
	lfu.freqMap[key] = 1

	return
}
//循环比较，取出最小频率的key值  todo 应该还可以加一个访问时间，频率相同情况下，判断访问时间
func (lfu *LFUCache)getMinFreqKey() interface{} {
	var minFreq uint16 = 0
	var tmp interface{}
	for key,freq := range lfu.freqMap{
		if minFreq == 0{
			minFreq, tmp = freq, key
			continue
		}
		if freq < minFreq {
			minFreq, tmp = freq, key
		}
	}
	return tmp
}

//打印一下
func (lfu *LFUCache)Print()  {
	lfu.list.Print()
}