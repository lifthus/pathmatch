package pathmatch

import "strings"

func splitFirstSegment(sep, path string) (seg string, remain string) {
	path = strings.Trim(path, sep)
	nextDelIdx := strings.Index(path, sep)
	if nextDelIdx == -1 {
		return path, ""
	}
	seg = path[:nextDelIdx]
	remain = path[nextDelIdx:]
	return seg, remain
}

func trimSep(str, sep string) string {
	str = strings.TrimPrefix(str, sep)
	return strings.TrimSuffix(str, sep)
}
