package iterations

import "strings"

func Repeat(term string, it int) string {
	var res strings.Builder
	for i := 0; i < it; i++ {
		res.WriteString(term)
	}
	return res.String()
}
