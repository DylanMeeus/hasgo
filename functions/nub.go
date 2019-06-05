package functions

func (s SliceType) Nub() (out SliceType) {
	if len(s) == 0 {
		return
	}

	singles := make(map[ElementType]int)
	for _, v := range s {
		singles[v] = 1
	}

	for k := range singles {
		out = append(out, k)
	}

	return
}
