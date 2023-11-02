package pathmatch

import "testing"

func TestFindOrCreateNodeForPath(t *testing.T) {
	IntMatchTree := &matchTree[int]{
		target: 1,
		ok:     true,
		next: matchMap[int]{
			"a": &matchTree[int]{
				next: matchMap[int]{
					"b": &matchTree[int]{
						target: 3,
						ok:     true,
						next:   matchMap[int]{},
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
