package functions

// Foldr reduces a list by iteratively applying f from right -> left. Thus, for an empty slice, the result is the default zero-value.
func (s SliceType) Foldr1(f func(e1, e2 ElementType) ElementType) (out ElementType) {
	if len(s) == 0 {
		return
	}
	s = s.Reverse()
	out = s[0]
	for _, v := range s[1:] {
		out = f(out, v)
	}
	return
}
