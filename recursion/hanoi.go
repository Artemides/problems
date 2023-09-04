package recursion

import "fmt"

func HanoiTowers(s, d, e string, n int) {
	if n <= 0 {
		return
	}

	fmt.Printf("Disk-%v from %s to %s\n", n, s, d)
	HanoiTowers(s, e, d, n-1)
	HanoiTowers(e, d, s, n-1)

}
