package lists_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tesserical/gtl/container/lists"
	"github.com/tesserical/gtl/predicate"
)

func TestList_General(t *testing.T) {
	ls := lists.NewList[int]()
	ls.PushBack(2)
	ls.PushBack(3)
	ls.PushBack(5)
	ls.PushFront(1)
	ls.Insert(ls.Size()-2, 4)
	assert.Equal(t, 5, ls.Size())
	assert.Equal(t, 1, ls.Front())
	assert.Equal(t, 5, ls.Back())
	ls.PopBack()
	assert.Equal(t, 4, ls.Size())
	assert.Equal(t, 1, ls.Front())
	assert.Equal(t, 4, ls.Back())
	assert.EqualValues(t, []int{1, 2, 3, 4}, ls.Slice())

	ls.Erase(1)
	assert.EqualValues(t, []int{1, 3, 4}, ls.Slice())

	count := 0
	for item := range ls.Begin() {
		assert.Equal(t, ls.Get(count), item)
		count++
	}

	count = ls.Size() - 1
	for item := range ls.End() {
		assert.Equal(t, ls.Get(count), item)
		count--
	}

	assert.True(t, ls.Contains(3, predicate.CompareFunc[int]))
	assert.False(t, ls.Contains(7, predicate.CompareFunc[int]))

	ls.Splice(0, lists.NewList[int](5, 6, 7))
	t.Log(ls.Slice())
}
