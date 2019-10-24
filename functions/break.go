package functions

// Break returns a tuple of any elements that do not satisfy the predicate up until the first time it passes, followed
// by the rest of the elements.
// Can be generated on any type.
func (s SliceType) Break(f func(ElementType) bool) (before SliceType, after SliceType) {
	if f == nil {
		return before, s
	}

	passed := false

	for _, v := range s {
		if passed || f(v) {
			after = append(after, v)
			passed = true
		} else {
			before = append(before, v)
		}
	}

	return
}
