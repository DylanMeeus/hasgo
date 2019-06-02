package functions

import "sort"

func (s SliceType) Mode() (out SliceType) {
	counts := make(map[ElementType]int)

	for _, v := range s {
		counts[v]++
	}

	type kv struct {
		Key   ElementType
		Value int
	}
	var ss []kv
	for k, v := range counts {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[j].Value > ss[i].Value
	})

	lastCount := -1
	for _, v := range ss {
		if v.Value < lastCount {
			return
		}
		out = append(out, v.Key)
		lastCount = v.Value
	}

	return
}
