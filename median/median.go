package median

import (
	"fmt"
	"math"
	"sort"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arrs := append(nums1, nums2...)
	sort.Ints(arrs)
	if long := len(arrs); long%2 == 1 {
		fmt.Println("1")
		return float64(arrs[long/2])
	} else {

		down := int(math.Floor(float64(long) / 2))
		fmt.Println(down)
		return (float64(arrs[down-1]) + float64(arrs[down])) / 2
	}
}
