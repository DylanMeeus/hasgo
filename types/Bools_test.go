package types

import (
	"testing"
)

var (
	boolsTests = []struct {
		input Bools
		and   bool
		or    bool
	}{
		{
			nil,
			true,
			true,
		},
		{
			Bools{true, true},
			true,
			true,
		},
		{
			Bools{},
			true,
			true,
		},
		{
			Bools{false},
			false,
			false,
		},
		{
			Bools{false, true, false},
			false,
			true,
		},
	}
)

func Test_BoolsAnd(t *testing.T) {
	for _, test := range boolsTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.And(); res != test.and {
				t.Errorf("expected %v but got %v", test.and, res)
			}
		})
	}
}

func Test_BoolsOr(t *testing.T) {
	for _, test := range boolsTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Or(); res != test.or {
				t.Errorf("expected %v but got %v", test.or, res)

			}
		})
	}
}
