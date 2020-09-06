package main

import (
	"sort"
	//"fmt"
)


type Freq struct {
	ele int 
	freq int
}



func findLeastNumOfUniqueInts(arr []int, k int) int {
	m := make(map[int]int)
	for  _,e := range arr {
		m[e] = m[e] + 1
	}
	// take a map , create array 
	a  := make([]Freq, len(m))
	idx := 0
	for ele,value := range m {
		a[idx].ele = ele
		a[idx].freq = value
		idx++
	}

	//fmt.Println("frquency arr ",a)
	sort.SliceStable(a, func(i, j int) bool {
		return a[i].freq < a[j].freq
	})

	//fmt.Println("frquency arr after sorting ",a)
	idx = 0
    for i , obj := range a {
		if k > obj.freq {
			k -= obj.freq
		} else if k == obj.freq{
			k = 0
			idx = i + 1
			break
		} else {
			k = 0
			idx = i 
			break 
		}
	}

	least := 0 
	least = len(a) - idx
	//fmt.Printf(" least value is %d",least)
	return least



    
}