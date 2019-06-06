package functions

func (s SliceType) Intersperse(value ElementType) (out SliceType) {
	for i, el := range s {
		out = append(out, el)
		if i == len(s)-1 {
			break
		}
		out = append(out, value)
	}
	return
}
