package functions

// Returns true if any of the elements satisfy the predicate
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
