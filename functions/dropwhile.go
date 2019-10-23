package functions

// DropWhile returns a new slice containing all elements after the predicate fails for the first time.
// Can be generated for any type.
func (s SliceType) DropWhile(f func(ElementType) bool) (out SliceType) {
	if f == nil {
		return s
	}
	failed := false
	for _, v := range s {
		if failed {
			out = append(out, v)
			continue
		}
		if !f(v) {
			out = append(out, v)
			failed = true
		}
	}
	return
}
