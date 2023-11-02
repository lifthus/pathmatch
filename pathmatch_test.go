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
