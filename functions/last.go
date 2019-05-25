package functions

// Returns the last element in the slice
// If no element is found, returns the zero-value of the type
func (s SliceType) Last() (out ElementType) {
	if len(s) > 0 {
		out = s[len(s)-1]
	}
	return
}
