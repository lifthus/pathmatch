package pathmatch

import "strings"

func splitFirstSegment(sep, path string) (seg string, remain string) {
	path = trimSep(path, sep)
	nextDelIdx := strings.Index(path, sep)
	if nextDelIdx == -1 {
		return path, ""
	}
	seg = path[:nextDelIdx]
	remain = path[nextDelIdx:]
	return seg, remain
}

func trimSep(str, sep string) string {
	for strings.HasPrefix(str, sep) {
		str = strings.TrimPrefix(str, sep)
	}
	for strings.HasSuffix(str, sep) {
		str = strings.TrimSuffix(str, sep)
	}
	return str
}
