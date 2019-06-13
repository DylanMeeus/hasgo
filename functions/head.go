package functions

// Returns the first element in the slice.
// If no element is found, returns the zero-value of the type.
// Can be generated for any type.
func (s SliceType) Head() (out ElementType) {
	if len(s) > 0 {
		out = s[0]
	}
	return
}
