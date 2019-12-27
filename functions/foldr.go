package functions

// Foldr reduces a list by iteratively applying f from right -> left. Thus, for an empty slice, the result is the default zero-value.
func (s SliceType) Foldr(e ElementType, f func(e1, e2 ElementType) ElementType) (out ElementType) {
	if len(s) == 0 {
		return
	}

	end := len(s) - 1
	out = f(s[end], e)

	for i := end - 1; i >= 0; i-- {
		out = f(s[i], out)
	}

	return
}
