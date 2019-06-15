package types

// Int is a wrapper around int64
//go:generate hasgo -T=int64 -S=Ints
type Int int64

// Ints is a wrapper around []int64
type Ints []int64

// IntReplicate creates a slice with `value` repeated `count` times
func IntReplicate(count uint64, value int64) (out Ints) {
	out = make(Ints, count)
	for i := uint64(0); i < count; i++ {
		out[i] = value
	}
	return
}

// IntRange creates a range similar to list comprehension in Haskell
// Ranges in Haskell are _inclusive_ on both bounds.
// IntRange(0,10) == [0..10] == []int{0,1,2,3,4,5,6,7,8,9,10}
func IntRange(start, stop int64) Ints {
	if stop-start <= 0 {
		return Ints{}
	}
	out := make(Ints, (stop-start)+1)
	var i int
	for start <= stop {
		out[i] = start
		i++
		start++
	}
	return out
}

// EqualsOrdered verifies that the elements in the slices are equal (same position)
func (is Ints) EqualsOrdered(other Ints) bool {
	if len(is) != len(other) {
		return false
	}
	for i := range is {
		if is[i] != other[i] {
			return false
		}
	}
	return true
}

// Equals verifies that both slices have the same elements, ignoring their position.
func (is Ints) Equals(other Ints) bool {
	if len(is) != len(other) {
		return false
	}
	contains := make(map[int64]struct{}, 0)
	for _, i := range is {
		contains[i] = struct{}{}
	}
	for _, i := range other {
		if _, ok := contains[i]; !ok {
			return false
		}
	}
	return true
}
