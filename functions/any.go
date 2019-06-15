package functions

// Any returns true if any of the elements satisfy the predicate.
// Can be generated for any type.
func (s SliceType) Any(f func(ElementType) bool) bool {
	if f == nil {
		return false
	}
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}
