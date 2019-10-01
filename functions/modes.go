package functions

// Modes returns the elements with highest frequency in the slice.
// Can be generated for any type.
func (s SliceType) Modes() (out SliceType) {
	if len(s) == 0 {
		return
	}

	counts := make(map[ElementType]int)
	for _, v := range s {
		counts[v]++
	}

	var max int
	for _, v := range counts {
		if v > max {
			max = v
		}
	}

	for k, v := range counts {
		if v == max {
			out = append(out, k)
		}
	}

	return
}
