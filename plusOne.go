package main 

import (
	"fmt"
)

func plusOne(digits []int) []int {
	n := len(digits)
	digits[n-1] += 1
	n--
	for ;n>0 && digits[n] >= 10;n-- {
		digits[n] %= 10;
		digits[n-1] += 1
	}
	if n == 0 && digits[n] >= 10 {
		digits[n] %= 10
		return append([]int{1}, digits...)
	} else {
		return digits
	}
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{1, 2, 9}))
	fmt.Println(plusOne([]int{9, 9, 9}))
}

