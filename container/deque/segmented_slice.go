package deque

import "iter"

const blockSize = 64 // number of elements per block

// A SegmentedSlice (double-ended queue) implemented as a segmented slice (map-of-blocks),
// inspired by the SGI C++ STL deque design. This gives O(1) push/pop at both ends,
// true O(1) random access, and good cache locality.
type SegmentedSlice[T any] struct {
	blocks     [][]T // slice of blocks
	headBlock  int   // index of block containing the first element
	headOffset int   // offset within headBlock for first element
	tailBlock  int   // index of block containing the last element
	tailOffset int   // offset within tailBlock for last element
	length     int   // total number of elements
}

// compile-time check
var _ Deque[any] = (*SegmentedSlice[any])(nil)

// NewSegmentedSlice creates an empty SegmentedSlice.
func NewSegmentedSlice[T any]() *SegmentedSlice[T] {
	// allocate one initial block
	b := make([][]T, 1)
	b[0] = make([]T, blockSize)
	return &SegmentedSlice[T]{
		blocks:     b,
		headBlock:  0,
		headOffset: blockSize / 2, // start in middle to allow growth both ways
		tailBlock:  0,
		tailOffset: blockSize/2 - 1,
		length:     0,
	}
}

// Len returns the number of elements.
func (d *SegmentedSlice[T]) Len() int {
	return d.length
}

// At returns element at index i (0 <= i < Len()), or false if out of range.
func (d *SegmentedSlice[T]) At(i int) (T, bool) {
	var zero T
	if i < 0 || i >= d.length {
		return zero, false
	}
	// compute absolute position from head
	abs := d.headOffset + i
	blockIdx := d.headBlock + abs/blockSize
	off := abs % blockSize
	return d.blocks[blockIdx][off], true
}

// PushBack inserts v at the end.
func (d *SegmentedSlice[T]) PushBack(v T) {
	// need new block?
	if d.tailOffset+1 == blockSize {
		// append block
		d.blocks = append(d.blocks, make([]T, blockSize))
		d.tailBlock++
		d.tailOffset = 0
	} else {
		d.tailOffset++
	}
	d.blocks[d.tailBlock][d.tailOffset] = v
	d.length++
}

// PushFront inserts v at the front.
func (d *SegmentedSlice[T]) PushFront(v T) {
	// need new block?
	if d.headOffset == 0 {
		// prepend block
		newBlk := make([]T, blockSize)
		d.blocks = append([][]T{newBlk}, d.blocks...)
		d.headBlock = 0
		d.tailBlock++
		d.headOffset = blockSize - 1
	} else {
		d.headOffset--
	}
	d.blocks[d.headBlock][d.headOffset] = v
	d.length++
}

// PopBack removes and returns the last element, or false if empty.
func (d *SegmentedSlice[T]) PopBack() (T, bool) {
	var zero T
	if d.length == 0 {
		return zero, false
	}
	v := d.blocks[d.tailBlock][d.tailOffset]
	// move tail inward
	d.tailOffset--
	d.length--
	// drop block if empty
	if d.tailOffset < 0 && d.tailBlock > d.headBlock {
		d.blocks = d.blocks[:len(d.blocks)-1]
		d.tailBlock--
		d.tailOffset = blockSize - 1
	}
	return v, true
}

// PopFront removes and returns the first element, or false if empty.
func (d *SegmentedSlice[T]) PopFront() (T, bool) {
	var zero T
	if d.length == 0 {
		return zero, false
	}
	v := d.blocks[d.headBlock][d.headOffset]
	// move head inward
	d.headOffset++
	d.length--
	// drop block if empty
	if d.headOffset == blockSize && d.headBlock < d.tailBlock {
		d.blocks = d.blocks[1:]
		d.headOffset = 0
		d.tailBlock--
		d.headBlock = 0
	}
	return v, true
}

// Begin returns a sequence of elements from the beginning to the end.
func (d *SegmentedSlice[T]) Begin() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := 0; i < d.length; i++ {
			v, ok := d.At(i)
			if !ok {
				break
			}
			if !yield(v) {
				break
			}
		}
	}
}

// End returns a sequence of elements from the end to the beginning.
func (d *SegmentedSlice[T]) End() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := d.length - 1; i >= 0; i-- {
			v, ok := d.At(i)
			if !ok {
				break
			}
			if !yield(v) {
				break
			}
		}
	}
}
