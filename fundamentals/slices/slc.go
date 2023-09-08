package slices

import "fmt"

func Run() {
	slc := []int{1, 2, 3, 4, 5}
	fmt.Printf("address: %p \n", &slc)
	slc = append(slc, 6)
	fmt.Printf("address 2: %p \n", &slc)
	fmt.Printf("len: %d cap: %d", len(slc), cap(slc))
}

func Seasons() {
	months := [...]string{"", "january", "february", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december"}
	q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(q2, summer)
	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)

}

func reverse(slc []int) {
	for i, j := 0, len(slc)-1; i < j; i, j = i+1, j-1 {
		slc[i], slc[j] = slc[j], slc[i]
	}
}

func Rotate(times int) {
	ints := []int{1, 2, 3, 4, 5, 6}
	reverse(ints[:times])
	fmt.Println(ints)
	reverse(ints[times:])
	fmt.Println(ints)
	reverse(ints)
	fmt.Println(ints)
}
