package pathmatch

import (
	"fmt"
	"strings"
)

type PathTargetMap[T any] map[string]T

// NewPathMatcher generates a new Matcher with "/" as sep.
// PathTargetMap is a map that contains path string as key and target T as value.
func NewPathMatcher[T any](pathTargetMap PathTargetMap[T]) (*Matcher[T], error) {
	return NewMatcher[T]("/", pathTargetMap)
}

// NewMatcher generates a new Matcher with given sep.
// PathTargetMap is a map that contains path string as key and target T as value.
func NewMatcher[T any](sep string, pathTargetMap PathTargetMap[T]) (*Matcher[T], error) {
	mm, err := buildMatchTree[T](sep, pathTargetMap)
	if err != nil {
		return nil, err
	}
	return &Matcher[T]{
		sep:     sep,
		rootSeg: mm,
	}, nil
}

type Matcher[T any] struct {
	sep     string
	rootSeg *segNode[T]
}

func (mch *Matcher[T]) Match(path string) (target T, targetPath string, ok bool) {
	full := path
	cursn := mch.rootSeg
	for {
		seg, remain := splitFirstSegment(mch.sep, full)
		nextsn, ok := cursn.nextSegMap[seg]
		if !ok {
			// going back to root if there is no target.
			// but root may have no target as well.
			if !cursn.ok {
				cursn = mch.rootSeg
				full = path
			}
			if !strings.HasPrefix(full, mch.sep) {
				full = mch.sep + full
			}
			return cursn.target, full, cursn.ok
		}
		cursn = nextsn
		full = remain
	}
}

func newSegNode[T any]() *segNode[T] {
	return &segNode[T]{
		nextSegMap: map[string]*segNode[T]{},
	}
}

type segNode[T any] struct {
	target     T
	ok         bool
	nextSegMap map[string]*segNode[T]
}

func (sn *segNode[T]) setTarget(target T) error {
	if sn.ok {
		return fmt.Errorf("target has been already set")
	}
	sn.ok = true
	sn.target = target
	return nil
}
