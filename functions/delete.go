package functions

// delete the first element in the list.
func (s SliceType) Delete(e ElementType) (out SliceType) {
	deleted := false
	for _, v := range s {
		if deleted || v != e {
			out = append(out, v)
		} else {
			deleted = true
		}
	}
	return
}
