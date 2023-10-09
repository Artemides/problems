package testing

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("anitalavalatina") {
		t.Error(`IsPalindrome(anitalavalatina) = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome(kayak) = false`)
	}

	if !IsPalindrome("été") {
		t.Errorf("IsPalindrome(%q) =f alse", "été")
	}

	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf("IsPalindrome(%q) = false", input)
	}
}

func TestNotPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome(palindrome) = true`)
	}
}

func TestIspalindromeV2(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},
	}
	for _, test := range testCases {
		if got := IsPalindromeV2(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}
