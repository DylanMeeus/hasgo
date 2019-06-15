package functions

// Minimum returns the minimum of a slice.
// Can be generated for number types.
func (s SliceType) Minimum() ElementType {
	if len(s) == 0 {
		return 0 // bit strange?
	}
	min := s[0]
	for _, i := range s {
		if i < min {
			min = i
		}
	}
	return min
}
