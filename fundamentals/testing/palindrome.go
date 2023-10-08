package testing

func IsPalindrome(word string) bool {
	for idx, _rune := range word {
		if _rune != rune(word[len(word)-1-idx]) {
			return false
		}
	}

	return true
}
