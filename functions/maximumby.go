package functions

// MaximumBy returns the maximum elements according to a custom comparator.
// Can be generated for any type.
func (s SliceType) MaximumBy(f func(e1, e2 ElementType) ElementType) (max ElementType) {
	if len(s) == 0 {
		return
	}
	max = s[0]
	for _, el := range s[1:] {
		max = f(max, el)
	}
	return
}
