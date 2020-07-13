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
	obj := Constructor(4)
	param_1 := obj.Get(1)
	log.Println(param_1)
	obj.Put(1, 11)
	for k := obj.l.head; k != nil; k = k.next {
		log.Println(k, k.next)
	}
	obj.Put(2, 22)
	for k := obj.l.head; k != nil; k = k.next {
		log.Println(k, k.next)
	}
	obj.Put(3, 33)

	obj.Put(4, 44)
	for k := obj.l.head; k != nil; k = k.next {
		log.Println(k, k.next)
	}
	obj.Put(4, 44)

	log.Println("33l:", obj.l.tail)
	obj.Put(4, 44)
	log.Println("44l:", obj.l.tail)
	obj.Put(5, 55)
	log.Println("55l:", obj.l.tail)
	obj.Put(3, 33)

	for k := obj.l.head; k != nil; k = k.next {
		log.Println(k.key, k.val)
	}
	log.Println(obj)
}

func Constructor(capacity int) LRUCache {
	var lru LRUCache
	var ln ListNode
	var l List
	l.head = &ln
	l.tail = &ln
	lru.m = make(map[int]int)
	lru.s = capacity
	lru.l = &l
	return lru
}

func (l *List) addFirst(key, val int) {
	var ln ListNode
	ln.key = key
	ln.val = val
	if l.size == 0 {
		l.head = &ln
		l.tail = &ln
	} else {
		ln.next = l.head
		l.head.prev = &ln
		l.head = &ln
	}
	l.size++
}

func (l *List) removeLast() int {
	if l.size == 0 {
		return -1
	}
	tail := l.tail
	key := tail.key
	if l.size == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.tail.prev.next = nil
		l.tail = l.tail.prev
	}
	l.size--
	return key
}

func (l *List) remove(key int) {
	for node := l.head; node != nil; node = node.next {
		if node.key == key {
			if node.prev != nil {
				node.prev.next = node.next
			}
			if node.next != nil {
				node.next.prev = node.prev
			}
			node = nil
			l.size--
			break
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
	}
	this.m[key] = value
	this.l.addFirst(key, value)
}
