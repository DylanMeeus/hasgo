package functions

// Tails returns all tails of a sequence, in order of large to small, as if it were called recursively.
// Can be generated for any type.
func (s SliceType) Tails() (out SliceSliceType) {
	slicecopy := append([]ElementType(nil), s...)
	for range s {
		out = append(out, slicecopy)
		slicecopy = slicecopy[1:]
	}
	out = append(out, make(SliceType, 0))
	return
}
