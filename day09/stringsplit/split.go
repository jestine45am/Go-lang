package stringsplit

import "strings"
// example Split("abc", "b") -> ["a", "c"]
func Split(str, sep string) []string {
	var ret []string
	var n int
	for {
		n = strings.Index(str, sep)
		if n == -1 {
			ret = append(ret, str)
			break
		}
		ret = append(ret, str[:n])
		str = str[n+len(sep):]
	}
	return ret
}

