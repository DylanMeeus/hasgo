package functions

// Elem returns true if the slice contains the element
// Can be generated for any type.
func (s SliceType) Elem(el ElementType) bool {
	for _, e := range s {
		if e == el {
			return true
		}
	}
	return false
}
