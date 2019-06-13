package functions

// Returns the average of all elements in the slice.
// Can be generated for all number types.
func (s SliceType) Average() float64 {
	var sum ElementType
	if len(s) == 0 {
		return float64(0)
	}
	for _, i := range s {
		sum += i
	}
	return float64(sum) / float64(len(s))
}
