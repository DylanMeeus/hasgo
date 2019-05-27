package functions

// tests if the slice is empty
func (s SliceType) Null() bool {
	return len(s) == 0
}
