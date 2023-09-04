package hackerrank

import "fmt"

func SockMerchant(n int32, ar []int32) int32 {
	pairs := make(map[int32]int)

	matches := int32(0)
	for _, sock := range ar {
		pairs[sock] += 1
		if pairs[sock]%2 == 1 {
			matches += 1
		}
	}
	fmt.Printf("%+v\n", pairs)
	fmt.Println(n, matches, (n - matches))
	return (n - matches)
}
