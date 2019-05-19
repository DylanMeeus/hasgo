package functions

func (s SliceType) Filter(f func(ElementType) bool) (out SliceType) {
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return
}
