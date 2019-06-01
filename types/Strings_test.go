package types

import (
	"strings"
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

	stringsPredicateTests = []struct {
		input  Strings
		predicate func(s string) bool
		any bool
		all bool
	}{
		{
			Strings{"a", "b", "c"},
			func(s string) bool { return s == "b" },
			true,
			false,
		},
		{
			Strings{"b", "b", "b"},
			func(s string) bool { return s == "b" },
			true,
			true,
		},
		{
			Strings{"b", "b", "b"},
			func(s string) bool { return s == "c" },
			false,
			false,
		},
	}

	stringsTakeTests = []struct {
		input      Strings
		takeAmount uint64
		take       Strings
		init       Strings
		tail       Strings
		head       string
		last       string
	}{
		{
			Strings{"a", "b", "c"},
			2,
			Strings{"a", "b"},
			Strings{"a", "b"},
			Strings{"b", "c"},
			"a",
			"c",
		},
	}

	stringsReverseTests = []struct {
		input  Strings
		output Strings
	}{
		{
			Strings{"a", "b", "c"},
			Strings{"c", "b", "a"},
		},
	}

	stringsUnconsTests = []struct {
		input Strings
		head  string
		tail  Strings
	}{
		{
			Strings{"a", "b", "c"},
			"a",
			Strings{"b", "c"},
		},
	}

	stringsMapTests = []struct {
		input   Strings
		mapfunc func(string) string
		output  Strings
	}{
		{
			Strings{"a", "b", "c"},
			func(s string) string { return strings.ToUpper(s) },
			Strings{"A", "B", "C"},
		},
	}

	stringsSortTests = []struct {
		input  Strings
		output Strings
	}{
		{
			Strings{"d", "a", "b", "b"},
			Strings{"a", "b", "b", "d"},
		},
	}

	stringsLengthTests = []struct {
		input  Strings
		output int
	}{
		{
			Strings{"a", "b", "c"},
			3,
		},
		{
			Strings{},
			0,
		},
		{
			nil,
			0,
		},
	}

	stringsBoolTests = []struct {
		input Strings
		null  bool
	}{
		{
			Strings{"a"},
			false,
		},
		{
			Strings{},
			true,
		},
		{
			nil,
			true,
		},
	}

	stringsIntercalateTests = []struct {
		input       Strings
		intercalate [][]string
		output      Strings
	}{
		{
			Strings{", "},
			[][]string{
				{"dylan", "meeus"},
				{"ana", "esparza"},
			},
			Strings{"dylan", "meeus", ", ", "ana", "esparza"},
		},
		{
			Strings{", "},
			[][]string{
				{"dylan", "meeus"},
				{"ana", "esparza"},
				{"chris", "evans"},
			},
			Strings{"dylan", "meeus", ", ", "ana", "esparza", ", ", "chris", "evans"},
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

func Test_StringsTake(t *testing.T) {
	for _, test := range stringsTakeTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Take(test.takeAmount); !res.EqualsOrdered(test.take) {
				t.Errorf("expected %v but got %v", test.take, res)
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

func Test_StringsUncons(t *testing.T) {
	for _, test := range stringsUnconsTests {
		t.Run("", func(t *testing.T) {
			if head, tail := test.input.Uncons(); head != test.head || !tail.EqualsOrdered(test.tail) {
				t.Errorf("expected (%v,%v) but got (%v,%v)", test.head, test.tail, head, tail)
			}
		})
	}
}

func Test_StringsMap(t *testing.T) {
	for _, test := range stringsMapTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Map(test.mapfunc); !res.EqualsOrdered(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_StringsMapOriginalUnchanged(t *testing.T) {
	for _, test := range stringsMapTests {
		t.Run("", func(t *testing.T) {
			if test.input.Map(test.mapfunc); test.input.EqualsOrdered(test.output) {
				t.Errorf("expected %v but got %v", test.input, test.output)
			}
		})
	}
}

func Test_StringsSort(t *testing.T) {
	for _, test := range stringsSortTests {
		t.Run("", func(t *testing.T) {
			if test.input.Sort(); test.input.EqualsOrdered(test.output) {
				t.Errorf("expected %v but got %v", test.input, test.output)
			}
		})
	}
}

func Test_StringsLength(t *testing.T) {
	for _, test := range stringsLengthTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Length(); res != test.output {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}



func Test_StringsAll(t *testing.T) {
	for _, test := range stringsPredicateTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.All(test.predicate); res != test.all {
				t.Errorf("expected %v but got %v", test.all, res)
			}
		})
	}
}

func Test_StringsAny(t *testing.T) {
	for _, test := range stringsPredicateTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Any(test.predicate); res != test.any {
				t.Errorf("expected %v but got %v", test.any, res)
			}
		})
	}
}

func Test_StringsNull(t *testing.T) {
	for _, test := range stringsBoolTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Null(); res != test.null {
				t.Errorf("expected %v but got %v", test.null, res)
			}
		})
	}
}

func Test_StringsIntercalate(t *testing.T) {
	for _, test := range stringsIntercalateTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Intercalate(test.intercalate); !res.EqualsOrdered(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}
