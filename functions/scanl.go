package functions

// Scanl reduces a list by iteratively applying f from left->right and then returns each iteration in a slice.
func (s SliceType) Scanl(e ElementType, f func(e1, e2 ElementType) ElementType) (out SliceType) {
	if len(s) == 0 {
		return
	}

	out = append(out, e)
	last := e

	for _, v := range s {
		last = f(last, v)
		out = append(out, last)
	}

	return
}
