package functions

func (s SliceType) Nub() (out SliceType) {
	if len(s) == 0 {
		return
	}

	contains := make(map[ElementType]struct{})
	for _, v := range s {
		if _, ok := contains[v]; !ok {
			contains[v] = struct{}{}
			out = append(out, v)
		}
	}
	return
}
