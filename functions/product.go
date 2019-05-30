package functions

func (s SliceType) Product() ElementType {
	var prod ElementType
	for _, v := range s {
		prod += v
	}
	return prod
}
