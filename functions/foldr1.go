package functions

// Foldr1 reduces a list by iteratively applying f from right -> left. Thus, for an empty slice, the result is the default zero-value.
func (s SliceType) Foldr1(f func(e1, e2 ElementType) ElementType) (out ElementType) {
	if len(s) == 0 {
		return
	}

	end := len(s) - 1
	out = s[end]

	for i := end - 1; i >= 0; i-- {
		out = f(s[i], out)
	}

	return
}
