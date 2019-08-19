package functions

// Foldr reduces a list by iteratively applying f from left->right, starting at the default
// zero-value of an element. Thus, for an empty slice, the result is the default zero-value.
func (s SliceType) Foldr(f func(e1, e2 ElementType) ElementType) (out ElementType) {
	for _, v := range s {
		out = f(out, v)
	}
	return
}
