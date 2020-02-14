package utils

import (
	"fmt"
	"strings"
)

func MapToString(m map[string]int) string {
	// Find the largest length key
	biggestLength := 0
	for k, _ := range m {
		if len(k) > biggestLength {
			biggestLength = len(k)
		}
	}

	// Print map
	var sb strings.Builder
	for k, v := range m {
		sb.WriteString(fmt.Sprintf("%-20v%d\n", k, v))
	}
	return sb.String()
}