package cluster

import "strings"

func escapeDouble(s string) string {
	if strings.Contains(s, `"`) {
		return "'" + s + "'"
	}
	return s
}
func escapeSingle(s string) string {
	if strings.Contains(s, `'`) {
		return "''" + s + "''"
	}
	return s
}
func escape(s string) string {
	return escapeSingle(escapeDouble(s))
}
