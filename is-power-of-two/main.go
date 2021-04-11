package main

import (
	"fmt"
)

func isPowerOfTwo(i int) bool {
	return i != 0 && (i & (i-1)) == 0
}

func main() {
	fmt.Println(isPowerOfTwo(0))
}
