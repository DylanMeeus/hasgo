package types

import (
	"testing"
)

// unit testing the Ints
var (
	intsSumTests = []struct {
		input  Ints
		output int64 
	}{
	{
		Ints([]int64{1,2,3}),
		6,
	},
	}
)

func Test_Sum(t *testing.T) {
	for _,test := range intsSumTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Sum(); res != test.output {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}
