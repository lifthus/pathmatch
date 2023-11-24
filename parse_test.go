package pathmatch

import "testing"

func TestSplitFirstSegmentOnlyWithSeps(t *testing.T) {
	QnA := map[string][]string{
		"":   {"", ""},
		"/":  {"", "/"},
		"//": {"", "//"},
	}
	for q, a := range QnA {
		seg, remain := splitFirstSegment(TestDel, q)
		if seg != a[0] || remain != a[1] {
			t.Errorf("%s is expected to be split into %s and %s, but got %s and %s", q, a[0], a[1], seg, remain)
		}
	}
}

func TestSplitFirstSegment(t *testing.T) {
	QnA := map[string][]string{
		"/a/b/c":   {"a", "/b/c"},
		"a/b/c":    {"a", "/b/c"},
		"a/b/c/":   {"a", "/b/c/"},
		"/a/b/c//": {"a", "/b/c//"},
		"a/b/c///": {"a", "/b/c///"},
	}

	for q, a := range QnA {
		seg, remain := splitFirstSegment(TestDel, q)
		if seg != a[0] || remain != a[1] {
			t.Errorf("%s is expected to be split into %s and %s, but got %s and %s", q, a[0], a[1], seg, remain)
		}
	}
}

func TestSplitFirstSegmentWithLongerSep(t *testing.T) {
	QnA := map[string][]string{
		"@#a@#b@#c":     {"a", "@#b@#c"},
		"a@#b@#c":       {"a", "@#b@#c"},
		"a@#b@#c@#":     {"a", "@#b@#c@#"},
		"@#a@#b@#c@#@#": {"a", "@#b@#c@#@#"},
		"a@#b@#c@#@#@":  {"a", "@#b@#c@#@#@"},
	}

	for q, a := range QnA {
		seg, remain := splitFirstSegment("@#", q)
		if seg != a[0] || remain != a[1] {
			t.Errorf("%s is expected to be split into %s and %s, but got %s and %s", q, a[0], a[1], seg, remain)
		}
	}
}

func TestTrimSepPrefix(t *testing.T) {
	QnA := map[string]string{
		"!@#a!@#b!@#c":       "a!@#b!@#c",
		"a!@#b!@#c":          "a!@#b!@#c",
		"a!@#b!@#c!@#":       "a!@#b!@#c!@#",
		"!@#a!@#b!@#c!@#!@#": "a!@#b!@#c!@#!@#",
		"a!@#b!@#c!@#!@#!@#": "a!@#b!@#c!@#!@#!@#",
		"!@#!@#a!@#b!@#!@#":  "a!@#b!@#!@#",
	}

	for q, a := range QnA {
		if str := trimSepPrefix(q, "!@#"); str != a {
			t.Errorf("%s is expected to be trimmed into %s, but got %s", q, a, str)
		}
	}
}
