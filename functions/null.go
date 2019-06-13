package functions

// Returns true the slice is empty.
// Can be generated for any type.
func (s SliceType) Null() bool {
	return len(s) == 0
}
