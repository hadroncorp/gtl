package lists

import (
	"iter"
)

type listNode[T any] struct {
	data T
	next *listNode[T]
	prev *listNode[T]
}

// List is a doubly linked list implementation.
// It allows for efficient insertion and removal of elements from both ends.
// It is not thread-safe and should be used in a single-threaded context.
type List[T any] struct {
	head  *listNode[T]
	tail  *listNode[T]
	count int
}

// NewList creates a new [List] instance.
func NewList[T any]() *List[T] {
	return &List[T]{}
}

// PushBack adds an element to the end of the list.
func (l *List[T]) PushBack(data T) {
	l.count++
	newNode := &listNode[T]{data: data}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	l.tail.next = newNode
	newNode.prev = l.tail
	l.tail = newNode
}

// PushFront adds an element to the front of the list.
func (l *List[T]) PushFront(data T) {
	l.count++
	newNode := &listNode[T]{data: data}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	newNode.next = l.head
	l.head.prev = newNode
	l.head = newNode
}

// Insert adds an element at the specified position in the list.
func (l *List[T]) Insert(pos int, data T) {
	if pos < 0 || pos >= l.count {
		return
	}
	if pos == 0 {
		l.PushFront(data)
		return
	}
	if pos == l.count-1 {
		l.PushBack(data)
		return
	}

	l.count++
	newNode := &listNode[T]{data: data}
	current := l.head
	for i := 0; i < pos; i++ {
		current = current.next
	}
	newNode.next = current.next
	newNode.prev = current
	current.next.prev = newNode
	current.next = newNode
}

// Front returns the first element of the list.
func (l *List[T]) Front() T {
	if l.head == nil {
		var zeroVal T
		return zeroVal
	}
	return l.head.data
}

// Back returns the last element of the list.
func (l *List[T]) Back() T {
	if l.tail == nil {
		var zeroVal T
		return zeroVal
	}
	return l.tail.data
}

// At returns the element at the specified position in the list.
func (l *List[T]) At(pos int) T {
	if pos < 0 || pos >= l.count {
		var zeroVal T
		return zeroVal
	}
	current := l.head
	for i := 0; i < pos; i++ {
		current = current.next
	}
	return current.data
}

// PopBack removes and returns the last element of the list.
func (l *List[T]) PopBack() T {
	if l.tail == nil {
		var zeroVal T
		return zeroVal
	}

	l.count--
	data := l.tail.data
	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	}
	if l.count == 0 {
		l.head = nil
	}
	return data
}

// PopFront removes and returns the first element of the list.
func (l *List[T]) PopFront() T {
	if l.head == nil {
		var zeroVal T
		return zeroVal
	}

	l.count--
	data := l.head.data
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	}
	if l.count == 0 {
		l.tail = nil
	}
	return data
}

// Remove removes the element at the specified position in the list.
func (l *List[T]) Remove(pos int) {
	if pos < 0 || pos >= l.count {
		return
	}
	if pos == 0 {
		l.PopFront()
		return
	}
	if pos == l.count-1 {
		l.PopBack()
		return
	}

	l.count--
	current := l.head
	for i := 0; i < pos; i++ {
		current = current.next
	}
	current.prev.next = current.next
	current.next.prev = current.prev
}

// Size returns the number of elements in the list.
func (l *List[T]) Size() int {
	return l.count
}

// IsEmpty returns true if the list is empty.
func (l *List[T]) IsEmpty() bool {
	return l.count == 0
}

// Clear removes all elements from the list.
func (l *List[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.count = 0
}

// Slice returns a slice containing all elements in the list.
func (l *List[T]) Slice() []T {
	slice := make([]T, 0, l.count)
	current := l.head
	for current != nil {
		slice = append(slice, current.data)
		current = current.next
	}
	return slice
}

// Begin returns a sequence that iterates over the list from the head to the tail.
func (l *List[T]) Begin() iter.Seq[T] {
	return func(yield func(T) bool) {
		current := l.head
		for current != nil {
			if !yield(current.data) {
				return
			}
			current = current.next
		}
	}
}

// End returns a sequence that iterates over the list from the tail to the head.
func (l *List[T]) End() iter.Seq[T] {
	return func(yield func(T) bool) {
		current := l.tail
		for current != nil {
			if !yield(current.data) {
				return
			}
			current = current.prev
		}
	}
}
