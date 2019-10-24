package functions

// Span returns a tuple of any elements that satisfy the predicate up until the first failure, followed by
// the rest of the elements.
// Can be generated for any type.
func (s SliceType) Span(f func(ElementType) bool) (before SliceType, after SliceType) {
	if f == nil {
		return before, s
	}

	failed := false

	for _, v := range s {
		if failed {
			after = append(after, v)
			continue
		}
		if !f(v) {
			after = append(after, v)
			failed = true
			continue
		}
		before = append(before, v)
	}

	return
}
