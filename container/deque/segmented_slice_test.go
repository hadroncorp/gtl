package deque_test

import (
	"testing"

	"github.com/hadroncorp/gtl/container/deque"
)

func TestDeque_General(t *testing.T) {
	d := deque.NewSegmentedSlice[int]()
	d.PushFront(4)
	d.PushBack(5)
	d.PushFront(3)
	d.PushBack(6)
	d.PushFront(2)
	d.PushBack(7)
	d.PushFront(1)
	d.PushBack(8)

	for item := range d.Begin() {
		t.Log(item)
	}

	t.Log("--- END ---")
	for item := range d.End() {
		t.Log(item)
	}

	d.PopFront()
	d.PopBack()

	t.Log("--- AFTER POP ---")
	for item := range d.Begin() {
		t.Log(item)
	}
}
