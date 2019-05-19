package functions

// Take n-1 elements from a slice, where n = len(list)
func (s SliceType) Init() (out SliceType) {
	slicecopy := append([]ElementType(nil), s...)
	return slicecopy[:len(s)-1]
}
