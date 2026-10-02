package main

import "fmt"

// Node reprezentálja a lista egyetlen elemét
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// LinkedList reprezentálja magát a listát
type LinkedList[T any] struct {
	Head *Node[T]
}

// Új elem hozzáadása a lista elejére (O(1))
func (l *LinkedList[T]) Prepend(val T) {
	newNode := &Node[T]{
		Value: val,
		Next:  l.Head,
	}
	l.Head = newNode
}

// Új elem hozzáadása a lista végére (O(n))
func (l *LinkedList[T]) Append(val T) {
	newNode := &Node[T]{Value: val}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// Lista elemeinek kiíratása
func (l *LinkedList[T]) Print() {
	current := l.Head
	for current != nil {
		fmt.Printf("%v -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	list := LinkedList[int]{}

	list.Append(10)
	list.Append(20)
	list.Prepend(5)

	list.Print() // Kimenet: 5 -> 10 -> 20 -> nil
}
