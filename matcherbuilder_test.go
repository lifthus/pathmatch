package pathmatch

import "testing"

var (
	PathIntMap = map[string]int{
		"/":      1,
		"a/a":    2,
		"/a/b/":  3,
		"/a/b/c": 4,
		"/a///c": 5,
	}
)

func TestBuildMatchTree(t *testing.T) {
	root, err := buildMatchTree[int](TestDel, PathIntMap)
	if err != nil {
		t.Errorf("buildMatchTree has failed: %v", err)
	}
	switch {
	case root.target != 1:
		fallthrough
	case root.nextSegMap["a"].nextSegMap["a"].target != 2:
		fallthrough
	case root.nextSegMap["a"].nextSegMap["b"].target != 3:
		fallthrough
	case root.nextSegMap["a"].nextSegMap["b"].nextSegMap["c"].target != 4:
		fallthrough
	case root.nextSegMap["a"].nextSegMap["c"].target != 5:
		t.Errorf("buildMatchTree has built invalid tree: %v", root)
	}
}

func TestFindOrCreateNodeForPath(t *testing.T) {
	IntMatchTree := &segNode[int]{
		target: 1,
		ok:     true,
		nextSegMap: map[string]*segNode[int]{
			"a": {
				nextSegMap: map[string]*segNode[int]{
					"b": {
						target:     3,
						ok:         true,
						nextSegMap: map[string]*segNode[int]{},
					},
				},
			},
		},
	}
	if mt := findOrCreateNodeForPath[int](IntMatchTree, "/", TestDel); mt.target != 1 {
		t.Errorf("/ should have target 1 but has %d", mt.target)
	}
	if mt := findOrCreateNodeForPath[int](IntMatchTree, "/a", TestDel); mt.ok {
		t.Errorf("/a should have no target but has %d", mt.target)
	}
	if mt := findOrCreateNodeForPath[int](IntMatchTree, "/a/b", TestDel); mt.target != 3 {
		t.Errorf("/a/b should have target 3 but has %d", mt.target)
	}
	if mt := findOrCreateNodeForPath[int](IntMatchTree, "/a/b/c", TestDel); mt == nil {
		t.Errorf("/a/b/c should create new node but got nil")
	}
}
