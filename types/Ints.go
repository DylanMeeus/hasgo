package types

//go:generate hasgo -T=int64 -S=Ints
type Int int64
type Ints []int64


// Create a range like you would with list comprehension in Haskell
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
