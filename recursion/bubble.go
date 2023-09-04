package recursion

import "fmt"

func BubbleSort(slc []int, n int) {
	if len(slc) == 0 {
		return
	}

	for idx, num := range slc[:len(slc)-1] {
		if num > slc[idx+1] {
			aux := num
			slc[idx] = slc[idx+1]
			slc[idx+1] = aux
		}
	}

	BubbleSort(slc[:len(slc)-1], n+1)

}

func MathTableOf(n, i int) {
	if i < 0 {
		return
	}

	MathTableOf(n, i-1)
	fmt.Printf("%v * %v = %v \n", n, i, n*i)
}
