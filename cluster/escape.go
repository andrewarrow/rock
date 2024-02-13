package cluster

import "strings"

func escape(s string) string {
	if strings.Contains(s, `"`) {
		return "'" + s + "'"
	}
	return s
}
