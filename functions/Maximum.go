package functions

func (s SliceType) Maximum() (out ElementType) {
	if len(s) == 0 {
		return
	}
	for _, i := range s {
		if i > out {
			out = i
		}
	}
	return
}
