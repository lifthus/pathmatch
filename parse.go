package pathmatch

import "strings"

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
