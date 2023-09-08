package composite

import (
	"crypto/sha256"
	"fmt"
)

func Composite() {
	// arr := [...]int{19: 1}
	// a := [...]int{1, 2}
	// b := [...]int{2, 1}

}

func CompareHash() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}
