package main

import (
	"log"
)

type ListNode struct {
	key, val   int
	prev, next *ListNode
}

type List struct {
	head, tail *ListNode
	size       int
}

type LRUCache struct {
	m map[int]int
	l *List
	s int
}

func main() {
	obj := Constructor(3)
	param_1 := obj.Get(1)
	log.Println(param_1)
	obj.Put(1, 11)
	log.Println(obj)
	obj.Put(2, 22)
	obj.Put(3, 33)
	log.Println(obj)
	obj.Put(4, 44)
	log.Println(obj)
}

func Constructor(capacity int) LRUCache {
	var lru LRUCache
	lru.m = make(map[int]int)
	lru.s = capacity
	lru.l.size = 0
	return lru
}

func (l *List) addFirst(key, val int) {
	var ln *ListNode
	ln.key = key
	ln.val = val
	if l.size == 0 {
		l.head = ln
		l.tail = ln
	} else {
		head := l.head
		ln.next = head.next
		head.next.prev = ln
		l.head = ln
	}
	l.size++
}

func (l *List) removeLast() int {
	var val int
	val = l.tail.val
	if l.size == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = l.tail.prev
	}
	l.size--
	return val
}

func (l *List) remove(key int) {
	node := l.head
	for ; node != nil; node = node.next {
		if node.key == key {
			node.prev.next = node.next
			node.next.prev = node.prev
		}
	}
}

func (this *LRUCache) Get(key int) int {
	val, ok := this.m[key]
	if !ok {
		return -1
	}
	this.Put(key, val)
	return val
}

func (this *LRUCache) Put(key int, value int) {
	_, ok := this.m[key]
	if ok {
		this.l.remove(key)
		delete(this.m, key)
	} else {
		if this.l.size == this.s {
			lastKey := this.l.removeLast()
			delete(this.m, lastKey)
		}
		this.m[key] = value
	}
	this.l.addFirst(key, value)

}
