package main

import (
	"codewars/maximumsubarraysum"
	"fmt"
)

func main() {
	r := maximumsubarraysum.MaximumSubarraySum([]int{-2, -1, -3, -4, -1, -2, -1, -5, -4})
	fmt.Println("Maximum Subarray Sum: ", r)
}
