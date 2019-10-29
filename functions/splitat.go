package functions

// SplitAt splits the slice at the given index, returning before and after as a tuple.
// Can be generated for any type.
func (s SliceType) SplitAt(i int) (before, after SliceType) {
	for k, v := range s {
		if k < i {
			before = append(before, v)
		} else {
			after = append(after, v)
		}
	}
	return
}
