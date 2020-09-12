package main

import (
	"fmt"
)
func canJump(nums []int) bool {
	jumpY := false
	if len(nums) == 1 {
		return true
	} 
	reached := len(nums) - 1
	for  i:= len(nums) -2 ; i >= 0  ; i-- {
	   jump := nums[i]
	   fmt.Printf("jump at %v is %v \n",i,jump)
	   if  jump  >= reached - i {
		   fmt.Printf("reached %v from %v \n",i,reached)
		   reached = i 
		   if reached == 0 {
			  jumpY = true
			  break
		   }
	   }
	}
	return jumpY
}