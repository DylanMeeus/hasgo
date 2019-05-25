package functions

func (s SliceType) Sum() ElementType {
	var sum ElementType
	for _, v := range s {
		sum += v
	}
	return sum
}
