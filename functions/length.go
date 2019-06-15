package functions

// Length returns the length (len) of a slice.
// Can be generated for any type.
func (s SliceType) Length() int {
	return len(s)
}
