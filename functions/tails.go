package functions

// Tails returns all tails of a sequence, in order of large to small, as if it were called recursively.
// Can be generated for any type.
func (s SliceType) Tails() (out SliceSliceType) {
	for i := range s {
		scopy := append([]ElementType(nil), s...)
		out = append(out, scopy[i:])
	}
	out = append(out, make(SliceType, 0))
	return
}
