package main 

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	ret := math.MinInt32
	currentSum := 0
	for _, v := range nums {
		if currentSum > 0 {
			if currentSum + v > ret {
				ret = currentSum + v
			}
		} else {
			if v > ret {
				ret = v
			}
		}
		if (currentSum + v > 0) {
			currentSum += v
		} else {
			currentSum = 0
		}
	}
	return ret
}

func main() {
	a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(a))
}

