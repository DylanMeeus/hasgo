package functions

// Take n-1 elements from a slice, where n = len(list).
// Can be generated for any type.
func (s SliceType) Init() (out SliceType) {
	if len(s) == 0 {
		return
	}
	slicecopy := append([]ElementType(nil), s...)
	return slicecopy[:len(s)-1]
}
