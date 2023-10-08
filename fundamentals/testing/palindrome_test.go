package testing

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("anitalavalatina") {
		t.Error(`IsPalindrome(anitalavalatina) = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome(kayak) = false`)
	}
}

func TestNotPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome(palindrome) = true`)
	}
}
