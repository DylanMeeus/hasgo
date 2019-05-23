package functions

// Return a new slice with the map operation applied to each element.
func (s SliceType) Map(f func(ElementType) ElementType) (out SliceType) {
	if f == nil {
		return s
	}
	for _, v := range s {
		out = append(out, f(v))
	}
	return
}
