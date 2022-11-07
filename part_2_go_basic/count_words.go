package part2

import "strings"

//CountWords in textS
func CountWords(text string) map[string]int32 {
	txt := strings.Fields(text)

	m := make(map[string]int32)

	for _, v := range txt {
		m[v]++
	}

	return m
}
