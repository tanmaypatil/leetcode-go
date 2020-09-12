package main

import "testing"


func TestJump(t *testing.T) {
	input := []int {2,3,1,1,4}
	output := canJump(input)
	expected := true
    if  output != expected {
       t.Errorf("output was incorrect, got: %v, want: %v.", output, expected)
    }
}

func TestJump1(t *testing.T) {
	input := []int {3,2,1,0,4}
	output := canJump(input)
	expected := false
    if  output != expected {
       t.Errorf("output was incorrect, got: %v, want: %v.", output, expected)
    }
}

func TestJump2(t *testing.T) {
	input := []int {0,4,4,4,4,4}
	output := canJump(input)
	expected := false
    if  output != expected {
       t.Errorf("output was incorrect, got: %v, want: %v.", output, expected)
    }
}
