package pathmatch

import (
	"fmt"
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

func (mtr *segNode[T]) setTarget(target T) error {
	if mtr.ok {
		return fmt.Errorf("target has been already set")
	}
	mtr.ok = true
	mtr.target = target
	return nil
}
