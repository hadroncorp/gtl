package lists_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hadroncorp/gtl/container/lists"
)

func TestList_General(t *testing.T) {
	ls := lists.NewList[int]()
	ls.PushBack(2)
	ls.PushBack(3)
	ls.PushBack(5)
	ls.PushFront(1)
	ls.Insert(ls.Len()-2, 4)
	assert.Equal(t, 5, ls.Len())
	assert.Equal(t, 1, ls.Front())
	assert.Equal(t, 5, ls.Back())
	ls.PopBack()
	assert.Equal(t, 4, ls.Len())
	assert.Equal(t, 1, ls.Front())
	assert.Equal(t, 4, ls.Back())
	assert.EqualValues(t, []int{1, 2, 3, 4}, ls.Slice())

	ls.Remove(1)
	assert.EqualValues(t, []int{1, 3, 4}, ls.Slice())

	count := 0
	for item := range ls.Begin() {
		assert.Equal(t, ls.At(count), item)
		count++
	}

	count = ls.Len() - 1
	for item := range ls.End() {
		assert.Equal(t, ls.At(count), item)
		count--
	}
}
