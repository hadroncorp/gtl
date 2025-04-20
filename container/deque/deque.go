package deque

import "iter"

// A Deque is a double-ended queue interface.
// It allows for efficient insertion and removal of elements from both ends.
type Deque[T any] interface {
	// PushBack inserts v at the end.
	PushBack(v T)
	// PushFront inserts v at the front.
	PushFront(v T)
	// PopBack removes and returns the last element, or false if empty.
	PopBack() (T, bool)
	// PopFront removes and returns the first element, or false if empty.
	PopFront() (T, bool)
	// Begin returns a sequence of elements from the beginning to the end.
	Begin() iter.Seq[T]
	// End returns a sequence of elements from the end to the beginning.
	End() iter.Seq[T]
}
