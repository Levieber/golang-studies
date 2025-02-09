package iteration

import "strings"

func Repeat(character string, numOfIterations int) string {
	var repeated strings.Builder
	for i := 0; i < numOfIterations; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
