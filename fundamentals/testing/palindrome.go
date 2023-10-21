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

func IsPalindromeV3(word string) bool {
	letters := make([]rune, 0, len(word))
	for _, _rune := range word {
		if !unicode.IsLetter(_rune) {
			continue
		}

		letters = append(letters, unicode.ToLower(_rune))
	}

	n := len(letters) / 2
	for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}

	return true
}
