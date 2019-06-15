package functions

// Sum returns the sum of all elements in the slice.
// Can be generated for any number type.
func (s SliceType) Sum() ElementType {
	var sum ElementType
	for _, v := range s {
		sum += v
	}
	return sum
}
