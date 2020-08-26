package main

import (
	"fmt"
	"sync"
)

type LRUcache struct {
	Head     *Node
	Tail     *Node
	Map      sync.Map
	Capacity int
	Length   int
}

type Node struct {
	LongURL  string
	ShortURL string
	Prev     *Node
	Next     *Node
}

func (this *LRUcache) init(cap int) {
	this.Length = 0
	this.Capacity = cap
}

func (this *LRUcache) find(longUrl string) string {
	if value, ok := this.Map.Load(longUrl); ok {
		node := value.(*Node)
		this.toHead(node)
		return node.ShortURL
	}
	return ""
}

func (this *LRUcache) add(longUrl string, shortUrl string) {
	newNode := &Node{LongURL: longUrl, ShortURL: shortUrl, Next: this.Head}
	if this.Length >= this.Capacity {
		tail := this.Tail
		this.Map.Delete(tail.LongURL)
		this.Tail = tail.Prev
		tail.free()
		this.Tail.Next = nil
		this.Length--
	} else if this.Length == 0 {
		this.Head = newNode
		this.Tail = this.Head
		this.Map.Store(longUrl, newNode)
		this.Length++
		return
	}
	this.Head.Prev = newNode
	this.Head = newNode
	this.Map.Store(longUrl, newNode)
	this.Length++
}

func (this *LRUcache) toHead(node *Node) {
	if node == this.Head {
		return
	}
	node.Prev.Next = node.Next
	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		this.Tail = node.Prev
	}
	node.Prev = nil
	node.Next = this.Head
	this.Head.Prev = node
	this.Head = node
}

func (this *Node) free() {
	this.ShortURL = ""
	this.LongURL = ""
	this.Prev = nil
	this.Next = nil
}

func (this *LRUcache) show() {
	fmt.Println("map:")
	this.Map.Range(func(key, value interface{}) bool {
		fmt.Println(key)
		return true
	})
	fmt.Println("node list:")
	curr := this.Head
	for curr != nil {
		fmt.Println(curr.LongURL, curr.ShortURL)
		curr = curr.Next
	}
	fmt.Println("from tail:")
	curr = this.Tail
	for curr != nil {
		fmt.Println(curr.LongURL, curr.ShortURL)
		curr = curr.Prev
	}
	fmt.Println()
	fmt.Println()
}
