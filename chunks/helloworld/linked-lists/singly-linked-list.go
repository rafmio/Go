package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	info interface{}
	next *Node
}

type List struct {
	head *Node
}

func (l *List) Insert(d interface{}) {
	list := &Node{info: d, next: nil}
	if l.head == nil {
		l.head = list
	} else {
		p := l.head
		for p.next != nil {
			p = p.next
		}
		p.next = list
	}
}

func Show(l *List) {
	p := l.head
	for p != nil {
		fmt.Println("-> %v ", p.info)
		p = p.next
	}
}

func main() {
	sl := List{}
	for i := 0; i < 5; i++ {
		sl.Insert(rand.Intn(100))
	}
	Show(&sl)
}
