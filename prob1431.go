package main

import (
	"sort"

)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	out := make([]bool, len(candies))
	candiesCopy := make([]int , len(candies))
	copy(candiesCopy,candies)
	sort.Ints(candiesCopy)
	maxVal := candiesCopy[len(candiesCopy) - 1 ]
	for i,s := range(candies) {
		if s + extraCandies >= maxVal {
           out[i] = true
		} else {
			out[i] = false
		}
	}
	return out

}
