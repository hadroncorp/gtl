package lists

import (
	"iter"

	"github.com/tesserical/gtl/container"
	"github.com/tesserical/gtl/predicate"
)

type doublyLinkedListNode[T any] struct {
	data T
	next *doublyLinkedListNode[T]
	prev *doublyLinkedListNode[T]
}

// DoublyLinkedList is a doubly linked list implementation.
// It allows for efficient insertion and removal of elements from both ends.
// It is not thread-safe and should be used in a single-threaded context.
type DoublyLinkedList[T any] struct {
	head  *doublyLinkedListNode[T]
	tail  *doublyLinkedListNode[T]
	count int
}

// Compile-time assertion
var _ List[any] = (*DoublyLinkedList[any])(nil)

// NewList creates a new [DoublyLinkedList] instance.
func NewList[T any](vals ...T) *DoublyLinkedList[T] {
	ls := &DoublyLinkedList[T]{}
	for i := range vals {
		ls.PushBack(vals[i])
	}
	return ls
}

// PushBack adds an element to the end of the list.
func (l *DoublyLinkedList[T]) PushBack(data T) {
	l.count++
	newNode := &doublyLinkedListNode[T]{data: data}
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
func (l *DoublyLinkedList[T]) PushFront(data T) {
	l.count++
	newNode := &doublyLinkedListNode[T]{data: data}
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
func (l *DoublyLinkedList[T]) Insert(pos int, data T) {
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
	newNode := &doublyLinkedListNode[T]{data: data}
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
func (l *DoublyLinkedList[T]) Front() T {
	if l.head == nil {
		var zeroVal T
		return zeroVal
	}
	return l.head.data
}

// Back returns the last element of the list.
func (l *DoublyLinkedList[T]) Back() T {
	if l.tail == nil {
		var zeroVal T
		return zeroVal
	}
	return l.tail.data
}

func (l *DoublyLinkedList[T]) Get(pos int) T {
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
func (l *DoublyLinkedList[T]) PopBack() T {
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
func (l *DoublyLinkedList[T]) PopFront() T {
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

func (l *DoublyLinkedList[T]) Erase(pos int) {
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

// Clear removes all elements from the list.
func (l *DoublyLinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.count = 0
}

// Slice returns a slice containing all elements in the list.
func (l *DoublyLinkedList[T]) Slice() []T {
	slice := make([]T, 0, l.count)
	current := l.head
	for current != nil {
		slice = append(slice, current.data)
		current = current.next
	}
	return slice
}

// Begin returns a sequence that iterates over the list from the head to the tail.
func (l *DoublyLinkedList[T]) Begin() iter.Seq[T] {
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
func (l *DoublyLinkedList[T]) End() iter.Seq[T] {
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

func (l *DoublyLinkedList[T]) Size() int {
	return l.count
}

func (l *DoublyLinkedList[T]) Empty() bool {
	return l.count == 0
}

func (l *DoublyLinkedList[T]) Contains(t T, pred predicate.Binary[T]) bool {
	for item := range l.Begin() {
		if pred(item, t) {
			return true
		}
	}
	return false
}

func (l *DoublyLinkedList[T]) Iterator() iter.Seq[T] {
	return l.Begin()
}

func (l *DoublyLinkedList[T]) Add(t T) {
	l.PushBack(t)
}

func (l *DoublyLinkedList[T]) AddAll(container container.Container[T]) {
	for item := range container.Iterator() {
		l.PushBack(item)
	}
}

func (l *DoublyLinkedList[T]) RemoveIf(pred predicate.Unary[T]) {
	i := 0
	for item := range l.Begin() {
		if pred(item) {
			l.Erase(i)
			return
		}
		i++
	}
}

func (l *DoublyLinkedList[T]) Sort(less predicate.Binary[T]) {
	// TODO implement me
	panic("implement me")
}

func (l *DoublyLinkedList[T]) Reverse(less predicate.Binary[T]) {
	// TODO implement me
	panic("implement me")
}

func (l *DoublyLinkedList[T]) Splice(pos int, other *DoublyLinkedList[T]) {
	if pos < 0 || pos >= l.count {
		return
	}

	if pos == 0 {
		oldHead := l.head
		l.head = other.head
		l.tail.next = oldHead
		return
	} else if pos == l.count-1 {
		l.PushBack(other.Back())
		return
	}
	current := l.head
	i := 0
	for current != nil {
		if i == pos {
			next := current.next
			current.next = other.head
			other.tail = next
			break
		}
		current = current.next
		i++
	}
}
