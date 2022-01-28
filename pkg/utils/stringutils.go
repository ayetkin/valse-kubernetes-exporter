package utils

import "strings"

// SplitKindName using to find kind from pod name
func SplitKindName(str string) string {
	split := strings.Split(str, "-")
	var splitArr []string
	for i := 0; i < len(split)-1; i++ {
		splitArr = append(splitArr, split[i])
	}
	return strings.Join(splitArr, "-")
}
