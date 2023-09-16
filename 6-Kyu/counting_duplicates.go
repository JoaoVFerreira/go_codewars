package _6Kyu

import (
	"strings"
)

func CountDuplicates(givenString string) int {
 charCount := make(map[string]int)
 givenString = strings.ToLower(givenString)
 for _, char := range givenString {
		charStr := string(char)
		 if _, exists := charCount[charStr]; exists {
				charCount[charStr]++
		 	} else {
				charCount[charStr] = 1
		}
 }
 duplicateCount := 0
 for _, count := range charCount {
		 if count > 1 {
			duplicateCount++
		}
 }
 return duplicateCount
}