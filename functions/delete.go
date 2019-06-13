package functions

// Delete the first element in the list.
// Can be generated for any type.
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
