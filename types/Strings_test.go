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
		input     Strings
		predicate func(s string) bool
		any       bool
		all       bool
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

	stringsReplicateTests = []struct {
		count  uint64
		value  string
		output Strings
	}{
		{5, "leet", Strings{"leet", "leet", "leet", "leet", "leet"}},
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

	stringsIntersperseTests = []struct {
		input       Strings
		intersperse string
		output      Strings
	}{
		{
			nil,
			"leet",
			Strings{},
		},
		{
			Strings{"dylan", "ana"},
			"meeus",
			Strings{"dylan", "meeus", "ana"},
		},
		{
			Strings{},
			"leet",
			Strings{},
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

	stringsModesTests = []struct {
		input  Strings
		output Strings
	}{
		{
			Strings{},
			Strings{},
		},
		{
			Strings{"dylan", "ana"},
			Strings{"dylan", "ana"},
		},
		{
			Strings{"dylan", "dylan", "ana", "ana"},
			Strings{"dylan", "ana"},
		},
		{
			Strings{"dylan", "dylan", "ana"},
			Strings{"dylan"},
		},
	}

	stringsNubTests = []struct {
		input  Strings
		output Strings
	}{
		{
			nil,
			Strings{},
		},
		{
			Strings{},
			Strings{},
		},
		{
			Strings{""},
			Strings{""},
		},
		{
			Strings{"dylan", "ana"},
			Strings{"dylan", "ana"},
		},
		{
			Strings{"dylan", "ana", "ana", "dylan"},
			Strings{"dylan", "ana"},
		},
		{
			Strings{"dylan", "dylan", "dylan"},
			Strings{"dylan"},
		},
	}

	// haskell string-specific operations
	stringsStringTests = []struct {
		input   Strings
		unlines string
		unwords string
	}{
		{
			nil,
			"",
			"",
		},
		{
			Strings{},
			"",
			"",
		},
		{
			Strings{"hello", "world"},
			"hello\nworld",
			"hello world",
		},
		{
			Strings{"Hello", "", "World"},
			"Hello\n\nWorld",
			"Hello  World",
		},
	}

	stringsDeleteTests = []struct {
		input   Strings
		element string
		output  Strings
	}{
		{
			nil,
			"a",
			Strings{},
		},
		{
			Strings{},
			"",
			Strings{},
		},
		{
			Strings{"a", "b", "a", "c"},
			"a",
			Strings{"b", "a", "c"},
		},
	}

	stringsWordsTests = []struct {
		input  string
		output Strings
	}{

		{
			"",
			Strings{},
		},
		{
			"hello world",
			Strings{"hello", "world"},
		},
		{
			"hello",
			Strings{"hello"},
		},
	}

	stringsLinesTests = []struct {
		input  string
		output Strings
	}{
		{
			"",
			Strings{},
		},
		{
			"hello",
			Strings{"hello"},
		},
		{
			"hello\nworld",
			Strings{"hello", "world"},
		},
	}

	stringsMinMaxByTests = []struct {
		input      Strings
		comparator func(s1, s2 string) string
		max        string
	}{
		{
			nil,
			func(s1, s2 string) string { return "" },
			"",
		},
		{
			Strings{"a", "bb", "aaa", "cc"},
			func(s1, s2 string) string {
				if len(s1) > len(s2) {
					return s1
				}
				return s2
			},
			"aaa",
		},
	}

	stringsFoldTests = []struct {
		input     Strings
		init      string
		foldrfunc func(s1, s2 string) string
		output    string
	}{
		{
			nil,
			"",
			func(s1, s2 string) string { return s1 },
			"",
		},
		{
			Strings{},
			"",
			func(s1, s2 string) string { return s1 },
			"",
		},
		{
			Strings{"one"},
			"zero",
			func(s1, s2 string) string {
				return s1 + " " + s2
			},
			"one zero",
		},
		{
			Strings{"one", "two", "three", "ten"},
			"",
			func(s1, s2 string) string {
				if len(s1) > len(s2) {
					return s1
				}
				return s2
			},
			"three",
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

func Test_StringsReplicate(t *testing.T) {
	for _, test := range stringsReplicateTests {
		t.Run("", func(t *testing.T) {
			if res := StringReplicate(test.count, test.value); !res.Equals(test.output) {
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

func Test_StringsIntersperse(t *testing.T) {
	for _, test := range stringsIntersperseTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Intersperse(test.intersperse); !res.Equals(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
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

func Test_StringsModes(t *testing.T) {
	for _, test := range stringsModesTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Modes(); !res.Equals(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_StringsNub(t *testing.T) {
	for _, test := range stringsNubTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Nub(); !res.Equals(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_StringsUnlines(t *testing.T) {
	for _, test := range stringsStringTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Unlines(); res != test.unlines {
				t.Errorf("expected %v but got %v", test.unlines, res)
			}
		})
	}
}

func Test_StringsUnwords(t *testing.T) {
	for _, test := range stringsStringTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Unwords(); res != test.unwords {
				t.Errorf("expected %v but got %v", test.unwords, res)
			}
		})
	}
}

func Test_StringsDelete(t *testing.T) {
	for _, test := range stringsDeleteTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Delete(test.element); !res.EqualsOrdered(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_StringsWords(t *testing.T) {
	for _, test := range stringsWordsTests {
		t.Run("", func(t *testing.T) {
			if res := Words(test.input); !res.EqualsOrdered(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_StringsFoldr(t *testing.T) {
	for _, test := range stringsFoldTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Foldr(test.init, test.foldrfunc); res != test.output {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_StringsLines(t *testing.T) {
	for _, test := range stringsLinesTests {
		t.Run("", func(t *testing.T) {
			if res := Lines(test.input); !res.EqualsOrdered(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_StringsMaximumBy(t *testing.T) {
	for _, test := range stringsMinMaxByTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.MaximumBy(test.comparator); res != test.max {
				t.Errorf("expected %v but got %v", test.max, res)
			}
		})
	}
}
