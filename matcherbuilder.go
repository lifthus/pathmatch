package pathmatch

func buildMatchTree[T any](del string, ptm PathTargetMap[T]) (*matchTree[T], error) {
	mt := newMatchTree[T]()
	for path, target := range ptm {
		curmt := mt
		targetNode := findOrCreateNodeForPath[T](curmt, path, del)
		if err := targetNode.SetTarget(target); err != nil {
			return nil, err
		}
	}
	return mt, nil
}

func findOrCreateNodeForPath[T any](curmt *matchTree[T], path string, del string) *matchTree[T] {
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
