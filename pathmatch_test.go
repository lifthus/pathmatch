package pathmatch

import "testing"

var (
	TestDel    = "/"
	PathIntMap = map[string]int{
		"/":      1,
		"a/a":    2,
		"/a/b/":  3,
		"/a/b/c": 4,
		"/a///c": 5,
	}
)

func TestBuildMatchMap(t *testing.T) {
	// root, err := buildMatchTree[int](TestDel, PathIntMap)
	// if err != nil {
	// 	t.Errorf("buildMatchTree has failed: %v", err)
	// }
	// switch {
	// case root.target != 1:
	// 	fallthrough
	// case root.next["a"].next["a"].target != 2:
	// 	fallthrough
	// case root.next["a"].next["b"].target != 3:
	// 	fallthrough
	// case root.next["a"].next["b"].next["c"].target != 4:
	// 	fallthrough
	// case root.next["a"].next["c"].target != 5:
	// 	t.Errorf("buildMatchTree has built invalid tree: %v", root)
	// }
}

func TestSearchOrGenerateNodeForPath(t *testing.T) {
	IntMatchTree := &matchTree[int]{
		target: 1,
		ok:     true,
		next: matchMap[int]{
			"a": &matchTree[int]{
				next: matchMap[int]{
					"b": &matchTree[int]{
						target: 3,
						ok:     true,
					},
				},
			},
		},
	}
	if mt := searchOrGenerateNodeForPath[int](IntMatchTree, "/", TestDel); mt.target != 1 {
		t.Errorf("/ should have target 1 but has %d", mt.target)
	}
	if mt := searchOrGenerateNodeForPath[int](IntMatchTree, "/a", TestDel); mt.ok {
		t.Errorf("/a should have no target but has %d", mt.target)
	}
	if mt := searchOrGenerateNodeForPath[int](IntMatchTree, "/a/b", TestDel); mt.target != 3 {
		t.Errorf("/a/b should have target 3 but has %d", mt.target)
	}
}

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
