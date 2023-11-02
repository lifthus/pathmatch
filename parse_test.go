package pathmatch

import "testing"

func TestSplitFirstSegment(t *testing.T) {
	QnA := map[string][]string{
		"/a/b/c":   {"a", "/b/c"},
		"a/b/c":    {"a", "/b/c"},
		"a/b/c/":   {"a", "/b/c"},
		"/a/b/c//": {"a", "/b/c"},
		"a/b/c///": {"a", "/b/c"},
	}

	for q, a := range QnA {
		seg, remain := splitFirstSegment(TestDel, q)
		if seg != a[0] || remain != a[1] {
			t.Errorf("%s is expected to be split into %s and %s, but got %s and %s", q, a[0], a[1], seg, remain)
		}
	}
}
