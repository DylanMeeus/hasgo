package functions

// inserts the method receiver slice into the function slice at each step
func (s SliceType) Intercalate(ss SliceSliceType) (out SliceType) {
	for i, slice := range ss {
		for _, el := range slice {
			out = append(out, el)
		}
		if i == len(ss)-1 {
			break
		}
		for _, el := range s {
			out = append(out, el)
		}
	}
	return out
}
