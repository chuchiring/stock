package main

import (
	"container/list"
	"log"
)

type item struct {
	value   int
	element *list.Element
}

//LRUCache 1
type LRUCache struct {
	capacity int
	count    int
	items    map[int]*item
	queue    *list.List
}

//Constructor 1
func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity}
}

//Get 1
func (cache *LRUCache) Get(key int) int {
	if cache.items == nil {
		return -1
	}

	if v, ok := cache.items[key]; ok {
		cache.queue.Remove(v.element)
		v.element = cache.queue.PushBack(key)
		return v.value
	}
	return -1
}

//Put 1
func (cache *LRUCache) Put(key int, value int) {

	if (cache.items) == nil {
		cache.items = make(map[int]*item)
		cache.queue = list.New()
	}

	if v, ok := cache.items[key]; ok {
		v.value = value
		cache.queue.Remove(v.element)
		v.element = cache.queue.PushBack(key)
	} else {
		if cache.count == cache.capacity {
			element := cache.queue.Front()
			delete(cache.items, element.Value.(int))
			cache.queue.Remove(element)

			v := &item{value: value}
			cache.items[key] = v
			v.element = cache.queue.PushBack(key)
		} else {
			v := &item{value: value}
			cache.items[key] = v
			v.element = cache.queue.PushBack(key)
			cache.count++
		}
	}

}

func main() {
	cache := &LRUCache{capacity: 2}

	cache.Put(2, 1)
	cache.Put(2, 2)
	if cache.Get(2) != 2 {
		log.Fatal("error")
	} // 返回  1
	cache.Put(1, 1)
	cache.Put(4, 1)
	if cache.Get(2) != -1 {
		log.Fatal("error: Get(2) != -1")
	} // 返回 -1 (未找到)

	//输入
	//["LRUCache","put","put","get","put","put","get"]
	//[[2],[2,1],[2,2],[2],[1,1],[4,1],[2]]
	//输出
	//[null,null,null,1,null,null,1]
	//预期结果
	//[null,null,null,2,null,null,-1]
}
