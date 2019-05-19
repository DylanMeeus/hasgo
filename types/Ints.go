package types

//go:generate hasgo -T=int64 -S=Ints
type Int int64
type Ints []int64

func IntRange(start, stop int64) Ints {
	if stop-start <= 0 {
		return Ints{}
	}
	out := make(Ints, stop-start)
	var i int
	for start < stop {
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
