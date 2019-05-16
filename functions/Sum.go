package functions

func (s Slice) int {
	var sum int
	for _,v := range s {
		sum += v
	}
	return sum
}
