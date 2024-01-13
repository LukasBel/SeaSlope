package handlers

import (
	"regexp"
	"strings"
)

func CleanText(input []string) []string {
	cleanedList := []string{}

	for _, v := range input {
		re := regexp.MustCompile(`\s+`)
		cleanedList = append(cleanedList, re.ReplaceAllString(strings.TrimSpace(v), " "))
	}

	return cleanedList
}
