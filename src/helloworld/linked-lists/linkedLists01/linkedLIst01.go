package main

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) insertAtHead(data int) {
	temp1 := &Node{data, nil}

	if l.head == nil {
		l.head = temp1
	} else {
		temp2 := l.head
		l.head = temp1
		temp1.next = temp2
	}
	l.length += 1
}

func (l *LinkedList) insertAtTail(data int) {
	temp1 := &Node{data, nil}

	if l.head == nil {
		l.head = temp1
	} else {
		temp2 := l.head
		for temp2.next != nil {
			temp2 = temp2.next
		}
		temp2.next = temp1
	}
	l.length += 1
}

func (l *LinkedList) insert(n, data int) {
	if n == 0 {
		l.insertAtHead(data)
	} else if n == l.length-1 {
		l.insertAtTail(data)
	} else {
		temp1 := &Node{data, nil}
		temp2 := l.head

		for i := 0; i < n-1; i++ {
			temp2 = temp2.next
		}
		temp1.next = temp2.next
		temp2.next = temp1
		l.length += 1
	}
}

func main() {
	list := LinkedList{nil, 0}
}
