package slices

import (
	"fmt"
	"unicode"
)

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

func reverse(slc *[]int) {
	slice := *slc
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func Rotate(ints *[]int, times int) {
	leftSlc := (*ints)[:times]
	reverse(&leftSlc)
	rightSlc := (*ints)[times:]
	reverse(&rightSlc)
	reverse(ints)
}

func RotateSinglePass(ints []int, times int) []int {
	ints = append(ints[times:], ints[:times]...)
	return ints
}

func Append(x []int, y int) []int {
	var z []int
	_len := len(z) + 1
	if _len <= cap(x) {
		z = x[:_len]
	} else {
		zcap := _len
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, _len, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func NonEmpty(strings []string) []string {
	i := 0
	for _, v := range strings {
		if v != "" {
			strings[i] = v
			i++
		}
	}
	return strings[:i]
}

func NonEmptyAppend(strings []string) []string {
	out := strings[:0]
	for _, str := range strings {
		if str != "" {
			out = append(out, str)
		}
	}
	return out
}

func RemoveAt(strings []string, i uint) []string {
	copy(strings[i:], strings[i+1:])
	return strings[:len(strings)-1]
}
func RemoveAtUnordered(strings []string, i uint) []string {
	strings[i] = strings[len(strings)-1]
	return strings[:len(strings)-1]
}

func RemoveAdjacentDuplicated(ints *[]int) {
	slc := *ints
	if len(slc) < 2 {
		return
	}

	rest := slc[1:]
	if slc[0] == slc[1] {
		copy(slc, slc[1:])
		rest = slc[:len(slc)-1]
		slc = rest
	}
	RemoveAdjacentDuplicated(&rest)
}

func SquashSpaces(bs *[]byte) {
	bytes := *bs
	isRun := false
	lastSpace := -1
	runStart := -1
	for i, _rune := range bytes {
		if unicode.IsSpace(rune(_rune)) {
			if i-1 == lastSpace && !isRun {

				runStart = lastSpace
				isRun = true
			}
			lastSpace = i
			continue
		}

		if isRun {
			copy(bytes[runStart:], bytes[lastSpace:])
			lastSpace = -1
			runStart = -1
		}
		isRun = false

	}

}
