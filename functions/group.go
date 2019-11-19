package functions

// Group returns a list of lists where each list contains only equal elements and the concatenation of the
// result is equal to the argument.
// Can be generated for any type.
func (s SliceType) Group() (out SliceSliceType) {
	current := SliceType{}

	for i, v := range s {
		current = append(current, v)
		if i == len(s)-1 || v != s[i+1] {
			out = append(out, current)
			current = SliceType{}
		}
	}

	return
}
