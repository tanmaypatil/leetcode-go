package main

import "testing"
import "reflect"

func TestDup(t *testing.T) {
	input := []int {4,3,2,7,8,2,3,1}
	output := findDuplicates(input)
	expected := []int {2,3}
    if  ! reflect.DeepEqual(output,expected) {
       t.Errorf("output was incorrect, got: %v, want: %v.", output, expected)
    }
}

func TestDup1(t *testing.T) {
	input := []int {1,1}
	output := findDuplicates(input)
	expected := []int {1}
    if  ! reflect.DeepEqual(output,expected) {
       t.Errorf("output was incorrect, got: %v, want: %v.", output, expected)
    }
}

func TestDup2(t *testing.T) {
	input := []int {1,2}
	output := findDuplicates(input)
	expected := []int {}
    if  ! reflect.DeepEqual(output,expected) {
       t.Errorf("output was incorrect, got: %v, want: %v.", output, expected)
    }
}