package pathmatch

func buildMatchTree[T any](sep string, ptm PathTargetMap[T]) (*segNode[T], error) {
	sn := newSegNode[T]()
	for path, target := range ptm {
		cursn := sn
		targetNode := findOrCreateNodeForPath[T](cursn, path, sep)
		if err := targetNode.setTarget(target); err != nil {
			return nil, err
		}
	}
	return sn, nil
}

func findOrCreateNodeForPath[T any](cursn *segNode[T], path string, sep string) *segNode[T] {
	remain := path
	var seg string
	for {
		seg, remain = splitFirstSegment(sep, remain)
		if seg == "" {
			return cursn
		}
		_, ok := cursn.nextSegMap[seg]
		if !ok {
			cursn.nextSegMap[seg] = newSegNode[T]()
		}
		cursn = cursn.nextSegMap[seg]
	}
}
