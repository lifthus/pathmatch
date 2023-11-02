package pathmatch

func buildMatchTree[T any](del string, ptm PathTargetMap[T]) (*segNode[T], error) {
	mt := newSegNode[T]()
	for path, target := range ptm {
		curmt := mt
		targetNode := findOrCreateNodeForPath[T](curmt, path, del)
		if err := targetNode.setTarget(target); err != nil {
			return nil, err
		}
	}
	return mt, nil
}

func findOrCreateNodeForPath[T any](curmt *segNode[T], path string, del string) *segNode[T] {
	remain := path
	var seg string
	for {
		seg, remain = splitFirstSegment(del, remain)
		if seg == "" {
			return curmt
		}
		_, ok := curmt.nextSegMap[seg]
		if !ok {
			curmt.nextSegMap[seg] = newSegNode[T]()
		}
		curmt = curmt.nextSegMap[seg]
	}
}
