package main

import "testing"

func TestProduct(t *testing.T) {
	input := []int {3,4,5,2}
    total := maxProduct(input)
    if total != 12 {
       t.Errorf("Product was incorrect, got: %d, want: %d.", total, 12)
    }
}

func TestProduct1(t *testing.T) {
	input := []int {1,5,4,5}
    total := maxProduct(input)
    if total != 16 {
       t.Errorf("Product was incorrect, got: %d, want: %d.", total, 16)
    }
}

func TestProduct2(t *testing.T) {
	input := []int {10,8}
    total := maxProduct(input)
    if total != 63 {
       t.Errorf("Product was incorrect, got: %d, want: %d.", total, 63)
    }
}

func TestProduct3(t *testing.T) {
	input := []int {20}
    total := maxProduct(input)
    if total != 19 {
       t.Errorf("Product was incorrect, got: %d, want: %d.", total, 19)
    }
}

func TestProduct4(t *testing.T) {
	input := []int {}
    total := maxProduct(input)
    if total != 0 {
       t.Errorf("Product was incorrect, got: %d, want: %d.", total, 0)
    }
}