package main

type LRUCache struct {
	capacity int
	cache    map[int]int
	keys     []int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]int),
		keys:     make([]int, 0),
	}
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.cache[key]; !ok {
		return -1
	} else {
		// 将 key 变为最近使用
		this.makeRecently(key)
		return val
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; ok {
		// 修改 key 的值
		this.cache[key] = value
		// 将 key 变为最近使用
		this.makeRecently(key)
		return
	}

	if len(this.cache) >= this.capacity {
		// 链表头部就是最久未使用的 key
		oldestKey := this.keys[0]
		this.keys = this.keys[1:]
		delete(this.cache, oldestKey)
	}
	// 将新的 key 添加链表尾部
	this.cache[key] = value
	this.keys = append(this.keys, key)
}

func (this *LRUCache) makeRecently(key int) {
	val := this.cache[key]
	// 删除 key，重新插入到队尾
	delete(this.cache, key)
	this.cache[key] = val
	// Move the key to the end to mark it as recently used.
	this.keys = append(this.keys, key)
	// Remove the old occurrence of the key.
	index := 0
	for i, k := range this.keys {
		if k == key {
			index = i
			break
		}
	}
	this.keys = append(this.keys[:index], this.keys[index+1:]...)
}
