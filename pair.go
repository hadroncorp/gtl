package gtl

// Pair is a generic struct that represents a key-value pair.
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}
