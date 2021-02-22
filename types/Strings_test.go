package types

import (
	"strings"
	"testing"
)

// unit testing the Strings
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
			Strings{},
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
			nil,
			func(s string) bool { return s == "b" },
			Strings{},
		},
		{
			Strings{},
			func(s string) bool { return s == "b" },
			Strings{},
		},
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
			nil,
			func(s string) bool { return s == "b" },
			false,
			true,
		},
		{
			Strings{},
			func(s string) bool { return s == "b" },
			false,
			true,
		},
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
		input       Strings
		takeAmount  uint64
		take        Strings
		init        Strings
		tail        Strings
		head        string
		last        string
		whilePred   func(string) bool
		whileOutput Strings
	}{
		{
			nil,
			0,
			Strings{},
			Strings{},
			Strings{},
			"",
			"",
			func(s string) bool { return len(s) < 2 },
			Strings{},
		},
		{
			Strings{},
			0,
			Strings{},
			Strings{},
			Strings{},
			"",
			"",
			func(s string) bool { return len(s) < 2 },
			Strings{},
		},
		{
			Strings{"a", "b", "c"},
			2,
			Strings{"a", "b"},
			Strings{"a", "b"},
			Strings{"b", "c"},
			"a",
			"c",
			func(s string) bool { return len(s) < 2 },
			Strings{"a", "b", "c"},
		},
	}

	stringsReverseTests = []struct {
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
			nil,
			"",
			Strings{},
		},
		{
			Strings{},
			"",
			Strings{},
		},
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
			nil,
			func(s string) string { return strings.ToUpper(s) },
			Strings{},
		},
		{
			Strings{},
			func(s string) string { return strings.ToUpper(s) },
			Strings{},
		},
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
			nil,
			Strings{},
		},
		{
			Strings{},
			Strings{},
		},
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
			nil,
			[][]string{
				{"dylan", "meeus"},
				{"ana", "esparza"},
			},
			Strings{"dylan", "meeus", "ana", "esparza"},
		},
		{
			Strings{},
			[][]string{
				{"dylan", "meeus"},
				{"ana", "esparza"},
			},
			Strings{"dylan", "meeus", "ana", "esparza"},
		},
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
			nil,
			Strings{},
		},
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
			Strings{},
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
		input    Strings
		init     string
		foldfunc func(s1, s2 string) string
		foldl    string
		foldl1   string
		foldr    string
		foldr1   string
	}{
		{
			nil,
			"",
			func(s1, s2 string) string { return s1 },
			"",
			"",
			"",
			"",
		},
		{
			Strings{},
			"",
			func(s1, s2 string) string { return s1 },
			"",
			"",
			"",
			"",
		},
		{
			Strings{"one"},
			"zero",
			func(s1, s2 string) string {
				return s1 + " " + s2
			},
			"zero one",
			"one",
			"one zero",
			"one",
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
			"three",
			"three",
			"three",
		},
	}

	stringsElemTests = []struct {
		input  Strings
		needle string
		output bool
	}{
		{
			nil,
			"",
			false,
		},
		{
			Strings{},
			"",
			false,
		},
		{
			Strings{"hello", "world"},
			"hello",
			true,
		},
		{
			Strings{"Hello", "World"},
			"",
			false,
		},
	}

	stringsDropTests = []struct {
		input  Strings
		drop   int
		output Strings
	}{
		{
			nil,
			0,
			Strings{},
		},
		{
			Strings{},
			0,
			Strings{},
		},
		{
			Strings{"Chris", "Evans", "Dylan", "Meeus"},
			5,
			Strings{},
		},
		{
			Strings{"Chris", "Evans", "Dylan", "Meeus"},
			1,
			Strings{"Evans", "Dylan", "Meeus"},
		},
		{
			Strings{"Chris", "Evans", "Dylan", "Meeus"},
			3,
			Strings{"Meeus"},
		},
	}

	stringsSpanTests = []struct {
		input    Strings
		spanfunc func(string) bool
		before   Strings
		after    Strings
	}{
		{
			nil,
			func(s string) bool { return len(s) < 3 },
			Strings{},
			Strings{},
		},
		{
			Strings{},
			func(s string) bool { return len(s) < 3 },
			Strings{},
			Strings{},
		},
		{
			Strings{"ABC", "ABCD", "ABCDE", "ABCDEF"},
			nil,
			Strings{},
			Strings{"ABC", "ABCD", "ABCDE", "ABCDEF"},
		},
		{
			Strings{"ABC", "ABCD", "ABCDE", "ABCDEF"},
			func(s string) bool { return len(s) < 9 },
			Strings{"ABC", "ABCD", "ABCDE", "ABCDEF"},
			Strings{},
		},
		{
			Strings{"ABC", "ABCD", "ABCDE", "ABCDEF"},
			func(s string) bool { return len(s) > 9 },
			Strings{},
			Strings{"ABC", "ABCD", "ABCDE", "ABCDEF"},
		},
		{
			Strings{"ABC", "ABCD", "ABCDE", "ABCDEF"},
			func(s string) bool { return len(s) < 5 },
			Strings{"ABC", "ABCD"},
			Strings{"ABCDE", "ABCDEF"},
		},
	}

	stringsDropWhileTests = []struct {
		input    Strings
		dropfunc func(string) bool
		output   Strings
	}{
		{
			nil,
			func(s string) bool { return len(s) < 4 },
			Strings{},
		},
		{
			Strings{},
			func(s string) bool { return len(s) < 4 },
			Strings{},
		},
		{
			Strings{"abc", "abcd", "abcdef", "abcdefg"},
			nil,
			Strings{"abc", "abcd", "abcdef", "abcdefg"},
		},
		{
			Strings{"a", "ab", "abc", "abcd", "abcdef", "abcdefg"},
			func(s string) bool { return len(s) < 4 },
			Strings{"abcd", "abcdef", "abcdefg"},
		},
		{
			Strings{"abc", "abcdef", "abc"},
			func(s string) bool { return len(s) < 4 },
			Strings{"abcdef", "abc"},
		},
	}

	stringsBreakTests = []struct {
		input     Strings
		breakfunc func(string) bool
		before    Strings
		after     Strings
	}{
		{
			nil,
			func(s string) bool { return len(s) < 3 },
			Strings{},
			Strings{},
		},
		{
			Strings{},
			func(s string) bool { return len(s) < 3 },
			Strings{},
			Strings{},
		},
		{
			Strings{"abc", "abcd", "abcde"},
			nil,
			Strings{},
			Strings{"abc", "abcd", "abcde"},
		},
		{
			Strings{"abc", "abcd", "abcde"},
			func(s string) bool { return len(s) < 100 },
			Strings{},
			Strings{"abc", "abcd", "abcde"},
		},
		{
			Strings{"abc", "abcd", "abcde"},
			func(s string) bool { return len(s) > 100 },
			Strings{"abc", "abcd", "abcde"},
			Strings{},
		},
		{
			Strings{"abc", "abcd", "abcde"},
			func(s string) bool { return len(s) > 4 },
			Strings{"abc", "abcd"},
			Strings{"abcde"},
		},
	}

	stringsInitsTests = []struct {
		input  Strings
		output []Strings
	}{
		{
			nil,
			[]Strings{},
		},
		{
			Strings{},
			[]Strings{},
		},
		{
			Strings{"aaa", "bbb", "ccc"},
			[]Strings{
				{},
				{"aaa"},
				{"aaa", "bbb"},
				{"aaa", "bbb", "ccc"},
			},
		},
	}

	stringsTailsTests = []struct {
		input  Strings
		output []Strings
	}{
		{
			nil,
			[]Strings{},
		},
		{
			Strings{},
			[]Strings{},
		},
		{
			Strings{"aaa", "bbb", "ccc"},
			[]Strings{
				{"aaa", "bbb", "ccc"},
				{"bbb", "ccc"},
				{"ccc"},
				{},
			},
		},
	}

	stringsSplitAtTests = []struct {
		input  Strings
		i      int
		before Strings
		after  Strings
	}{
		{
			nil,
			5,
			Strings{},
			Strings{},
		},
		{
			Strings{},
			5,
			Strings{},
			Strings{},
		},
		{
			Strings{"aaa", "bbb", "ccc", "ddd", "eee"},
			0,
			Strings{},
			Strings{"aaa", "bbb", "ccc", "ddd", "eee"},
		},
		{
			Strings{"aaa", "bbb", "ccc", "ddd", "eee"},
			-1,
			Strings{},
			Strings{"aaa", "bbb", "ccc", "ddd", "eee"},
		},
		{
			Strings{"aaa", "bbb", "ccc", "ddd", "eee"},
			10,
			Strings{"aaa", "bbb", "ccc", "ddd", "eee"},
			Strings{},
		},
		{
			Strings{"aaa", "bbb", "ccc", "ddd", "eee"},
			5,
			Strings{"aaa", "bbb", "ccc", "ddd", "eee"},
			Strings{},
		},
		{
			Strings{"aaa", "bbb", "ccc", "ddd", "eee"},
			3,
			Strings{"aaa", "bbb", "ccc"},
			Strings{"ddd", "eee"},
		},
	}

	stringsGroupTests = []struct {
		input  Strings
		output []Strings
	}{
		{
			nil,
			[]Strings{},
		},
		{
			Strings{},
			[]Strings{},
		},
		{
			Strings{"a", "b", "c"},
			[]Strings{{"a"}, {"b"}, {"c"}},
		},
		{
			Strings{"a", "a", "b", "c", "c", "c", "d"},
			[]Strings{{"a", "a"}, {"b"}, {"c", "c", "c"}, {"d"}},
		},
	}

	stringsScanTests = []struct {
		input    Strings
		init     string
		scanfunc func(string, string) string
		scanl    Strings
	}{
		{
			nil,
			"",
			func(s1, s2 string) string { return s1 + " " + s2 },
			Strings{},
		},
		{
			Strings{},
			"",
			func(s1, s2 string) string { return s1 + " " + s2 },
			Strings{},
		},
		{
			Strings{"World"},
			"Hello",
			func(s1, s2 string) string { return s1 + " " + s2 },
			Strings{"Hello", "Hello World"},
		},
		{
			Strings{"Hello", "World", "Goodbye"},
			"Thing",
			func(s1, s2 string) string { return s1 + " " + s2 },
			Strings{"Thing", "Thing Hello", "Thing Hello World", "Thing Hello World Goodbye"},
		},
	}

	stringsIsPrefixOfTests = []struct {
		input  Strings
		param1 Strings
		output bool
	}{
		{
			nil,
			Strings{},
			true,
		},
		{
			Strings{},
			Strings{},
			true,
		},
		{
			Strings{"a", "b", "cd"},
			Strings{"a", "b"},
			false,
		},
		{
			Strings{"a", "b"},
			Strings{"a", "b", "cd"},
			true,
		},
		{
			Strings{"a", "b"},
			Strings{"a", "b"},
			true,
		},
		{
			nil,
			Strings{"a", "b"},
			true,
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

func Test_StringsTakewhile(t *testing.T) {
	for _, test := range stringsTakeTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.TakeWhile(test.whilePred); !res.EqualsOrdered(test.whileOutput) {
				t.Errorf("expected %v but got %v", test.whileOutput, res)
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
			if test.input.Map(test.mapfunc); len(test.input) > 0 && test.input.EqualsOrdered(test.output) {
				t.Errorf("expected %v but got %v", test.input, test.output)
			}
		})
	}
}

func Test_StringsSort(t *testing.T) {
	for _, test := range stringsSortTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Sort(); !res.EqualsOrdered(test.output) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_StringsSortOriginalUnchanged(t *testing.T) {
	for _, test := range stringsSortTests {
		t.Run("", func(t *testing.T) {
			if test.input.Sort(); len(test.input) > 0 && test.input.EqualsOrdered(test.output) {
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

func Test_StringsFoldl1(t *testing.T) {
	for _, test := range stringsFoldTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Foldl1(test.foldfunc); res != test.foldl1 {
				t.Errorf("expected %v but got %v", test.foldl1, res)
			}
		})
	}
}

func Test_StringsFoldl(t *testing.T) {
	for _, test := range stringsFoldTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Foldl(test.init, test.foldfunc); res != test.foldl {
				t.Errorf("expected %v but got %v", test.foldl, res)
			}
		})
	}
}

func Test_StringsFoldr(t *testing.T) {
	for _, test := range stringsFoldTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Foldr(test.init, test.foldfunc); res != test.foldr {
				t.Errorf("expected %v but got %v", test.foldr, res)
			}
		})
	}
}

func Test_StringsFoldr1(t *testing.T) {
	for _, test := range stringsFoldTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Foldr1(test.foldfunc); res != test.foldr1 {
				t.Errorf("expected %v but got %v", test.foldr1, res)
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

func Test_StringsElem(t *testing.T) {
	for _, test := range stringsElemTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Elem(test.needle); res != test.output {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_StringsDrop(t *testing.T) {
	for _, test := range stringsDropTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Drop(test.drop); !test.output.Equals(res) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_StringsSpan(t *testing.T) {
	for _, test := range stringsSpanTests {
		t.Run("", func(t *testing.T) {
			before, after := test.input.Span(test.spanfunc)
			if !test.before.Equals(before) || !test.after.Equals(after) {
				t.Errorf("expected (%v, %v) but got (%v, %v)", test.before, test.after, before, after)
			}
		})
	}
}

func Test_StringsDropWhile(t *testing.T) {
	for _, test := range stringsDropWhileTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.DropWhile(test.dropfunc); !test.output.Equals(res) {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}

func Test_StringsBreak(t *testing.T) {
	for _, test := range stringsBreakTests {
		t.Run("", func(t *testing.T) {
			before, after := test.input.Break(test.breakfunc)
			if !test.before.Equals(before) || !test.after.Equals(after) {
				t.Errorf("expected %v, %v but got %v, %v", test.before, test.after, before, after)
			}
		})
	}
}

func Test_StringsInits(t *testing.T) {
	for _, test := range stringsInitsTests {
		t.Run("", func(t *testing.T) {
			res := test.input.Inits()
			for i, v := range test.output {
				if !v.EqualsOrdered(res[i]) {
					t.Errorf("expected %v but got %v", v, res[i])
				}
			}
		})
	}
}

func Test_StringsTails(t *testing.T) {
	for _, test := range stringsTailsTests {
		t.Run("", func(t *testing.T) {
			res := test.input.Tails()
			for i, v := range test.output {
				if !v.EqualsOrdered(res[i]) {
					t.Errorf("expected %v but got %v", v, res[i])
				}
			}
		})
	}
}

func Test_StringsSplitAt(t *testing.T) {
	for _, test := range stringsSplitAtTests {
		t.Run("", func(t *testing.T) {
			if before, after := test.input.SplitAt(test.i); !test.before.Equals(before) || !test.after.Equals(after) {
				t.Errorf("expected %v, %v but got %v, %v", test.before, test.after, before, after)
			}
		})
	}
}

func Test_StringsGroup(t *testing.T) {
	for _, test := range stringsGroupTests {
		t.Run("", func(t *testing.T) {
			res := test.input.Group()
			for i, v := range test.output {
				if !v.EqualsOrdered(res[i]) {
					t.Errorf("expected %v but got %v", v, res[i])
				}
			}
		})
	}
}

func Test_StringsScanl(t *testing.T) {
	for _, test := range stringsScanTests {
		t.Run("", func(t *testing.T) {
			if res := test.input.Scanl(test.init, test.scanfunc); !test.scanl.Equals(res) {
				t.Errorf("expected %v but got %v", test.scanl, res)
			}
		})
	}
}

func Test_StringsIsPrefixOf(t *testing.T) {
	for _, test := range stringsIsPrefixOfTests {
		t.Run("", func(t *testing.T) {
			res := test.input.IsPrefixOf(test.param1)
			if res != test.output {
				t.Errorf("expected %v but got %v", test.output, res)
			}
		})
	}
}
