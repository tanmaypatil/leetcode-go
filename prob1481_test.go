package main

import "testing"

func TestLeastInt(t *testing.T) {
	input := []int {5,5,4}
	k := 1
    output := findLeastNumOfUniqueInts(input, k)
    if output != 1 {
       t.Errorf("output was incorrect, got: %d, want: %d.", output, 1)
    }
}

func TestLeastInt1(t *testing.T) {
	input := []int {4,3,1,1,3,3,2}
	k := 3
    output := findLeastNumOfUniqueInts(input, k)
    if output != 2 {
       t.Errorf("output was incorrect, got: %d, want: %d.", output, 2)
    }
}

func TestLeastInt2(t *testing.T) {
	input := []int {1,1,1}
	k := 2
    output := findLeastNumOfUniqueInts(input, k)
    if output != 1 {
       t.Errorf("output was incorrect, got: %d, want: %d.", output, 1)
    }
}