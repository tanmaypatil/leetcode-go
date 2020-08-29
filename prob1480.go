package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")
}

func runningSum(nums []int) []int {
	outarr := make([]int, len(nums))
	for i, s := range nums {
		fmt.Println(i, s)
		outarr[i] += s 
	}
    return outarr
}

func Sum(x int, y int) int {
    return x + y + 1
}

