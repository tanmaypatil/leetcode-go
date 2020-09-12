package main

import (
	"container/list"
	"fmt"
)

func canReach(arr []int, start int) bool {
	if arr[start] == 0  {
		return true
	}
	queue := list.New()
	visited := make([]bool,len(arr))
	queue.PushBack(start)
	fmt.Println("queue is ",queue)
	position := start
	jump := arr[position]
	fmt.Printf(" jump at the start %v",jump)
	// check till queue has elements suggesting 
	// that there are elements to be checked.
	for ( queue.Len () > 0) {
		position = queue.Front().Value.(int)
		fmt.Printf("starting with position [%v]\n ",position)
		rightIndex := position + arr[position] 
		fmt.Printf("right index after jump [%v]\n ",rightIndex)
		// if element at the jumped position is zero , path exists 
		if rightIndex < len(arr) && arr[rightIndex] == 0 {
			fmt.Printf("path to zero exists [%v]\n",rightIndex)
			return true
		}
		if rightIndex < len(arr) && !visited[rightIndex]  {
			queue.PushBack(rightIndex)
		}
		leftIndex := position - arr[position] 
		fmt.Printf("left index after jump [%v]\n ",leftIndex)
		if leftIndex >=0 && arr[leftIndex] == 0 {
			fmt.Printf("path to zero exists [%v]\n",leftIndex)
			return true
		}
		if leftIndex >= 0 && ! visited[leftIndex] {
			queue.PushBack(leftIndex)
		}
		queue.Remove(queue.Front())
		visited[position] = true
		fmt.Println("current queue is ",queue)

	}

	return false
}
