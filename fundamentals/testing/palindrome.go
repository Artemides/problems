package testing

import "unicode"

func IsPalindrome(word string) bool {
	for idx, _rune := range word {
		if _rune != rune(word[len(word)-1-idx]) {
			return false
		}
	}

	return true
}

func IsPalindromeV2(word string) bool {
	var letters []rune
	for _, _rune := range word {
		if !unicode.IsLetter(_rune) {
			continue
		}

		letters = append(letters, _rune)
	}

	for idx, _rune := range letters {
		if _rune != letters[len(letters)-1-idx] {
			return false
		}
	}
	return true
}
