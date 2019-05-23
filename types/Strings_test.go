package types

import (
	"testing"
)

// unit testing the Ints
var (
	stringsSumTests = []struct {
		input  Strings
		output string
	}{
		{
			nil,
			"",
		},
		{
			Strings([]string{}),
			"",
		},
		{
			Strings([]string{"a", "b", "c"}),
			"abc",
		},
	}

	stringsFilterTests = []struct {
		input  Strings
		filter func(s string) bool
		output Strings
	}{
		{
			Strings{"a", "b", "c"},
			func(s string) bool { return s == "b" },
			Strings{"b"},
		},
	}

	stringsTakeTests = []struct {
		input Strings
		init  Strings
		tail  Strings
		head  string
		last  string
	}{
		{
			Strings{"a", "b", "c"},
			Strings{"a", "b"},
			Strings{"b", "c"},
			"a",
			"c",
		},
	}

	stringsReverseTests = []struct{
		input Strings
		output Strings
	}{
		{
			Strings{"a","b","c"},
			Strings{"c","b","a"},
		},
	}
)

func Test_StringsSum(t *testing.T) {
	for _, test := range stringsSumTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Sum(); res != test.output {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_StringsFilter(t *testing.T) {
	for _, test := range stringsFilterTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Filter(test.filter); !res.Equals(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_StringsInit(t *testing.T) {
	for _, test := range stringsTakeTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Init(); !res.EqualsOrdered(test.init) {
				t.Errorf("expected %v but got %v", test.init, res)
			}
		})
	}
}

func Test_StringsTail(t *testing.T) {
	for _, test := range stringsTakeTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Tail(); !res.EqualsOrdered(test.tail) {
				t.Errorf("expected %v but got %v", test.tail, res)
			}
		})
	}
}

func Test_StringsHead(t *testing.T) {
	for _, test := range stringsTakeTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Head(); res != test.head {
				t.Errorf("expected %v but got %v", test.head, res)
			}
		})
	}
}

func Test_StringsLast(t *testing.T) {
	for _, test := range stringsTakeTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Last(); res != test.last {
				t.Errorf("expected %v but got %v", test.last, res)
			}
		})
	}
}

func Test_StringsReverse(t *testing.T) {
	for _, test := range stringsReverseTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Reverse(); !res.EqualsOrdered(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}
