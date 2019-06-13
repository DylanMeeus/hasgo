package functions

// Take [1 -> n] elements from a slice, where n = len(list)
// Returns an empty slice if there are less than 2 elements in slice
// Can be generated for any type.
func (s SliceType) Tail() (out SliceType) {
	if len(s) <= 1 {
		return
	}
	slicecopy := append([]ElementType(nil), s...)
	return slicecopy[1:]
}
