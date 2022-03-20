package utils

import "strings"

func GetPoints(problemID string, verdict string, points *map[string]int) int {
	if verdict == "Accepted" {
		return (*points)[strings.ToUpper(problemID)]
	}

	return 0
}
