package functions

// Filter returns a slice containing only the elements that match the predicate.
// Can be generated for any type.
func (s SliceType) Elem(el ElementType) bool {
	for _, e := range s {
		if e == el {
			return true
		}
	}
	return false
}
