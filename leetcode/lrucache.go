package main

import "fmt"

/**
 * LRU 缓存淘汰策略
 * leetcode@146 https://leetcode-cn.com/problems/lru-cache/
 * 要求 获取&写入 时间复杂度都为 O(1)
 */

func main() {
	var cache = Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // 1
	cache.Put(3, 3)
	fmt.Println(cache.Get(2)) // -1
	cache.Put(4, 4)
	fmt.Println(cache.Get(1)) // -1
	fmt.Println(cache.Get(3)) // 3
	fmt.Println(cache.Get(4)) // 4
}

type CacheNode struct {
	prev  *CacheNode
	next  *CacheNode
	value int
	key   int
}

type LRUCache struct {
	capacity int // 缓存容量
	hash     map[int]*CacheNode
	head     *CacheNode
	tail     *CacheNode
}

func Constructor(capacity int) LRUCache {
	var head, tail = &CacheNode{}, &CacheNode{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		hash:     make(map[int]*CacheNode),
		head:     head,
		tail:     tail,
	}
}

func (l *LRUCache) InsertToHead(p *CacheNode) {
	// 插入头部
	p.next = l.head.next
	p.prev = l.head
	l.head.next = p
	p.next.prev = p
	// l.head.next.prev = p	// (⊙o⊙)…
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.hash[key]; ok {
		node.prev.next = node.next
		node.next.prev = node.prev
		l.InsertToHead(node)
		return node.value
	}

	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.hash[key]; ok {
		node.prev.next = node.next
		node.next.prev = node.prev
		l.InsertToHead(node)
		node.value = value
	} else {
		var newNode = &CacheNode{value: value, key: key}

		if len(l.hash) == l.capacity {
			// 删除最后一个节点
			var trashNode = l.tail.prev
			l.tail.prev.prev.next = l.tail
			l.tail.prev = trashNode.prev
			delete(l.hash, trashNode.key)
			trashNode = nil
		}

		l.InsertToHead(newNode)
		l.hash[key] = newNode
	}
}
