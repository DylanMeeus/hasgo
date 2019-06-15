package functions

// Reverse returns the reversed slice.
// Can be generated for any type.
func (s SliceType) Reverse() (out SliceType) {
	for i := len(s) - 1; i >= 0; i-- {
		out = append(out, s[i])
	}
	return
}
