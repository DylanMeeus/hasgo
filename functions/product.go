package functions

// Product returns the product of all elements in the slice.
// Can be generated for any number type.
func (s SliceType) Product() ElementType {
	var prod ElementType
	for _, v := range s {
		prod += v
	}
	return prod
}
