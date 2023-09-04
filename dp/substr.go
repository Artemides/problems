package dp

import "fmt"

func MaxHalvesSubstringLen(str string) int {
	strLen := len(str)
	maxLength := 0

	for i := 0; i < strLen; i++ {
		for j := i + 1; j < strLen; j += 2 {
			substring := str[i : j+1]
			fmt.Println(substring)
			if maxLength > len(substring) {
				continue
			}

			half := len(substring) / 2
			left := sum(substring[:half])
			right := sum(substring[half:])

			if left == right {
				maxLength = len(substring)
			}
		}
	}
	return maxLength
}

func sum(str string) int {
	sum := 0
	for _, num := range str {
		sum += int(num)
	}
	return sum
}
