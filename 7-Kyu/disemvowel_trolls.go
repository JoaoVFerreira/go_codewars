package _7Kyu

import "strings"

/*
	Receive a random string, should remove vowels and return a string without the vowels
*/

func Disemvowel(comment string) string {
	vowels := "aeiouAEIOU"
	var commentSanitized strings.Builder

	for _, char := range comment {
		if !strings.ContainsRune(vowels, char) {
			commentSanitized.WriteRune(char)
		}
	} 
	return commentSanitized.String()
}

