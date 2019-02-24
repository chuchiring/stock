package main

import (
	"log"
)

var curTime int

type item struct {
	key, value int
	lastAccess int
}

//LRUCache 1
type LRUCache struct {
	capacity int
	count    int

	internalTime int

	items []item
}

//Constructor 1
func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity}
}

//Get 1
func (cache *LRUCache) Get(key int) int {
	if len(cache.items) == 0 {
		return -1
	}

	for i, item := range cache.items {
		if item.key == key && item.value != -1 {
			// cache.items[i].readCount++
			cache.items[i].lastAccess = cache.internalTime
			cache.internalTime++
			return item.value
		}
	}

	return -1
}

//Put 1
func (cache *LRUCache) Put(key int, value int) {
	//init
	if (cache.items) == nil {
		cache.items = make([]item, cache.capacity)
		for i := 0; i < cache.capacity; i++ {
			cache.items[i].value = -1
		}
	}

	//put
	emptyindex := -1
	foundindex := -1

	lessReadIndex := -1
	lastAccessTime := cache.internalTime

	for i, item := range cache.items {

		//empty
		if emptyindex == -1 && item.value == -1 {
			emptyindex = i
			break
		}

		//foundindex
		if foundindex == -1 && item.key == key {
			foundindex = i
		}

		//lessusecount
		if item.value != -1 {
			if item.lastAccess < lastAccessTime {
				lessReadIndex = i
				lastAccessTime = item.lastAccess
			}
		}
	}

	if foundindex != -1 {
		cache.items[foundindex].value = value
		cache.items[foundindex].lastAccess = cache.internalTime
		cache.internalTime++
		// fmt.Println(cache.items)
		return
	}

	if emptyindex != -1 {
		cache.items[emptyindex].key = key
		cache.items[emptyindex].value = value
		cache.items[emptyindex].lastAccess = cache.internalTime
		cache.internalTime++
		// fmt.Println(cache.items)
		return
	}

	if lessReadIndex != -1 {
		cache.items[lessReadIndex].key = key
		cache.items[lessReadIndex].value = value
		cache.items[lessReadIndex].lastAccess = cache.internalTime
		cache.internalTime++
		// fmt.Println(cache.items)
	} else {
		log.Fatal("it is error")
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
