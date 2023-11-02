package pathmatch

import "testing"

var (
	TestDel = "/"
)

func TestMatch(t *testing.T) {
	PathIntMap := map[string]int{
		"/":      1,
		"a/a":    2,
		"/a/b/":  3,
		"/a/b/c": 4,
		"/a///c": 5,
	}
	matcher, err := NewPathMatcher[int](PathIntMap)
	if err != nil {
		t.Errorf("NewPathMatcher has failed: %v", err)
	}

	if target, path, ok := matcher.Match("a"); target != 1 || path != "/a" || !ok {
		t.Errorf("<a> should be matched to <1, /a> but got <%d, %s> with <%t>", target, path, ok)
	}
	if target, path, ok := matcher.Match("/"); target != 1 || path != "/" || !ok {
		t.Errorf("</> should be matched to <1, /> but got <%d, %s> with <%t>", target, path, ok)
	}
	if target, path, ok := matcher.Match("a/a"); target != 2 || path != "/" || !ok {
		t.Errorf("<a/a> should be matched to <2, /> but got <%d, %s> with <%t>", target, path, ok)
	}
	if target, path, ok := matcher.Match("a/a/cde"); target != 2 || path != "/cde" || !ok {
		t.Errorf("<a/a/cde> should be matched to <2, /cde> but got <%d, %s> with <%t>", target, path, ok)
	}
	if target, path, ok := matcher.Match("a/c/123"); target != 5 || path != "/123" || !ok {
		t.Errorf("<a/c/123> should be matched to <5, /123> but got <%d, %s> with <%t>", target, path, ok)
	}
}
