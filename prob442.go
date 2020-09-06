package main

import (
	"sort"
	// "fmt"
)

func findDuplicates(nums []int) []int {
	a := make([]int, 0)
	sort.Ints(nums)
	// fmt.Println("after sorting :", nums)

	dupFound := false
	prevVal := -1
	for _, val := range nums {
		if prevVal != val {
			dupFound = false
			prevVal = val
			continue
		} else if prevVal == val {
			if dupFound != true {
				a = append(a, val)
				// fmt.Println("dup found , current dup array :", a)
			}
			prevVal = val
			dupFound = true
		}
	}
	// fmt.Println("final dup array :", a)
	return a

}
