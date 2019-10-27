package functions

// Inits returns all inits of a sequence, in order of small to large, as if it were called recursively.
// Can be generated for any type.
func (s SliceType) Inits() (out SliceSliceType) {
	// {1, 2, 3} -> {{}, {1}, {1, 2}, {1, 2, 3}}
	out = append(out, make(SliceType, 0))
	for i := range s {
		new := make(SliceType, i+1)
		for n := 0; n <= i; n++ {
			new[n] = s[n]
		}
		out = append(out, new)
	}
	return
}
