package functions

// All returns true if all elements of the slice satisfy the predicate.
// Can be generated for any type.
func (s SliceType) All(f func(ElementType) bool) bool {
	if f == nil {
		return false
	}
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}
