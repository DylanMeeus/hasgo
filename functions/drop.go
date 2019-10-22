package functions

func (s SliceType) Drop(i int) (out SliceType) {
	for i < len(s) {
		out = append(out, s[i])
		i++
	}
	return
}
