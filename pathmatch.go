package pathmatch

import (
	"fmt"
	"strings"
)

// base path list with eact T ( target path ) => cut path and proper T

type PathTargetMap[T any] map[string]T

// NewPathMatcher generates a new Matcher with "/" as delimiter.
// PathTargetMap is a map that contains path string as key and target T as value.
func NewPathMatcher[T any](pathTargetMap PathTargetMap[T]) (*Matcher[T], error) {
	return NewMatcher[T]("/", pathTargetMap)
}

// NewMatcher generates a new Matcher with given delimiter.
// PathTargetMap is a map that contains path string as key and target T as value.
func NewMatcher[T any](delimiter string, pathTargetMap PathTargetMap[T]) (*Matcher[T], error) {
	mm, err := buildMatchTree[T](delimiter, pathTargetMap)
	if err != nil {
		return nil, err
	}
	return &Matcher[T]{
		del:   delimiter,
		match: mm,
	}, nil
}

func buildMatchTree[T any](del string, ptm PathTargetMap[T]) (*matchTree[T], error) {
	mt := newMatchTree[T]()
	for path, target := range ptm {
		curmt := mt
		targetNode := searchOrGenerateNodeForPath[T](curmt, path, del)
		if err := targetNode.SetTarget(target); err != nil {
			return nil, err
		}
	}
	return mt, nil
}

func searchOrGenerateNodeForPath[T any](curmt *matchTree[T], path string, del string) *matchTree[T] {
	remain := path
	var seg string
	for {
		seg, remain = splitFirstSegment(del, remain)
		if seg == "" {
			return curmt
		}
		_, ok := curmt.next[seg]
		if !ok {
			curmt.next[seg] = newMatchTree[T]()
		}
		curmt = curmt.next[seg]
	}
}

func splitFirstSegment(del, path string) (seg string, remain string) {
	path = strings.Trim(path, del)
	nextDelIdx := strings.Index(path, del)
	if nextDelIdx == -1 {
		return path, ""
	}
	seg = path[:nextDelIdx]
	remain = path[nextDelIdx:]
	return seg, remain
}

type Matcher[T any] struct {
	del   string
	match *matchTree[T]
}

func newMatchTree[T any]() *matchTree[T] {
	return &matchTree[T]{
		next: matchMap[T]{},
	}
}

type matchTree[T any] struct {
	target T
	ok     bool
	next   matchMap[T]
}

func (mtr *matchTree[T]) SetTarget(target T) error {
	if mtr.ok {
		return fmt.Errorf("target has been already set")
	}
	mtr.ok = true
	mtr.target = target
	return nil
}

type matchMap[T any] map[string]*matchTree[T]
