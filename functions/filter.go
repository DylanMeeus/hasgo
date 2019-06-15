package functions

// Filter returns a slice containing only the elements that match the predicate.
// Can be generated for any type.
func (s SliceType) Filter(f func(ElementType) bool) (out SliceType) {
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return
}
