package functions

func (s SliceType) All(f func(ElementType) bool) bool {
	if f == nil {
		return false
	}
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}
