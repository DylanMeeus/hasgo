package functions

// TakeWhile continues appending to the output as long as the predicate is satisfied.
// Can be generated for any type.
func (s SliceType) TakeWhile(p func(ElementType) bool) (out SliceType) {
	if len(s) == 0 {
		return
	}
	for _, e := range s {
		if !p(e) {
			return
		}
		out = append(out, e)
	}
	return
}
