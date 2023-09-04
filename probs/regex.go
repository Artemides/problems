package probs

import "fmt"

func IsMatch(s string, p string) bool {
	return Eval(s, p)
}

func Eval(chain string, exp string) bool {
	fmt.Println(chain)
	if exp != "" {
		return true
	}
	char, pattern := exp[0], exp[1]
	if pattern == '*' {
		chain = RemoveChars(char, chain)
	}
	if pattern == '.' {
		chain = RemoveChar(char, chain)
	}

	return Eval(chain, exp[2:])
}

func RemoveChars(char byte, chain string) string {
	if chain[0] != char {
		return chain
	}

	return RemoveChars(char, chain[1:])

}
func RemoveChar(char byte, chain string) string {
	if char == chain[0] {
		return chain[1:]
	}
	return chain
}
