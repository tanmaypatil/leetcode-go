package main

import "testing"


func TestQ(t *testing.T) {
	input := []int {4,2,3,0,3,1,2}
	exists := canReach(input,5)
	if exists != true {
		t.Errorf("incorrect answer , got :%v , expected :%v",exists,true)
	}
}

func TestQ1(t *testing.T) {
	input := []int {4,2,3,0,3,1,2}
	exists := canReach(input,0)
	if exists != true {
		t.Errorf("incorrect answer , got :%v , expected :%v",exists,true)
	}
}

func TestQ2(t *testing.T) {
	input := []int {3,0,2,1,2}
	exists := canReach(input,2)
	if exists != false {
		t.Errorf("incorrect answer , got :%v , expected :%v",exists,false)
	}

}

func TestQ3(t *testing.T) {
	input := []int {5,5,0,5,5}
	exists := canReach(input,2)
	if exists != true {
		t.Errorf("incorrect answer , got :%v , expected :%v",exists,false)
	}
}

func TestQ4(t *testing.T) {
	input := []int {5,5,0,5,5}
	exists := canReach(input,2)
	if exists != true {
		t.Errorf("incorrect answer , got :%v , expected :%v",exists,true)
	}
}