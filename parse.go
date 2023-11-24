package pathmatch

import "strings"

func splitFirstSegment(sep, path string) (seg string, remain string) {
	seg = trimSepPrefix(path, sep)
	if seg == "" {
		return "", path
	}
	nextDelIdx := strings.Index(seg, sep)
	if nextDelIdx == -1 {
		return seg, ""
	}
	return seg[:nextDelIdx], seg[nextDelIdx:]
}

func trimSepPrefix(str, sep string) string {
	for strings.HasPrefix(str, sep) {
		str = strings.TrimPrefix(str, sep)
	}
	return str
}
