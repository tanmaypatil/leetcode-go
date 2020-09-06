package main

import (
	"sort"
)

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	if len(nums) == 1 {
      return (nums[0] - 1)
	}
	return (nums[len(nums)-1] -1) * (nums[len(nums)-2] -1 )
    
}