package main

import "fmt"

type Node2[T any] struct {
	Value T
	Next  *Node2[T]
	Prev  *Node2[T]
}

type LinkedList2[T any] struct {
	Head *Node2[T]
}

func (l *LinkedList2[T]) AppendToTheFront(val T) {
	newNode := &Node2[T]{
		Value: val,
		Next:  l.Head,
		Prev:  nil,
	}

	if l.Head != nil {
		l.Head.Prev = newNode
	}

	l.Head = newNode
}

func (l *LinkedList2[T]) AppendToBack(val T) {
	newNode := &Node2[T]{
		Value: val,
		Next:  nil, // A legutolsó elemnek nincs rákövetkezője
		Prev:  nil,
	}

	// 1. Eset: A lista teljesen üres
	if l.Head == nil {
		l.Head = newNode
		return
	}

	// 2. Eset: Elmegyünk a legutolsó létező elemig
	current := l.Head
	for current.Next != nil {
		current = current.Next // Javítva: léptetés a következő elemre
	}

	// 3. Eset: Kétoldali kapcsolat kiépítése a végén
	newNode.Prev = current
	current.Next = newNode
}

func (l *LinkedList2[T]) Print() {
	current := l.Head
	for current != nil {
		fmt.Printf("%v -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}
