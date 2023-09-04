package leetcode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
)

func MaxProfit(prices []int) int {
	ss := LoadTest()
	profit := Calc(ss, 0)
	fmt.Println(profit)
	return profit
}

func Calc(prices []int, profit int) int {
	if len(prices) == 0 {
		return profit
	}

	max, idx := Find(prices, true)
	left := prices[:idx]
	right := prices[idx+1:]

	if max == 0 {
		return profit
	}

	if len(left) == 0 {
		return Calc(right, profit)
	}

	sort.Ints(left)
	min := left[0]

	if max-min > profit && max != min {
		profit = max - min
	}
	fmt.Println("right ", right)
	return Calc(right, profit)

}

func Find(slc []int, findMax bool) (found, idx int) {
	if len(slc) == 0 {
		return
	}

	found = slc[0]
	for i, num := range slc {
		if num > found == findMax {
			found = num
			idx = i
		}
	}
	fmt.Printf("found %v %v %v \n", findMax, found, idx)
	return

}
func LoadTest() []int {
	path := "./leetcode/profit.txt"
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return []int{}
	}

	var data []int

	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return []int{}
	}

	return data
}
