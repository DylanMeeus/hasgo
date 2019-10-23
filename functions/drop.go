package functions

// Drop returns a new Slice with the elements after the provided key.
// Can be generated for any type.
func (s SliceType) Drop(i int) (out SliceType) {
	for i < len(s) {
		out = append(out, s[i])
		i++
	}
	return
}
