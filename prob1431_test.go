package main

import "testing"
import (
	"reflect"
)

func TestCandies(t *testing.T) {
	input := []int {2,3,5,1,3}
	extra := 3
	out := [] bool{true,true,true,false,true}

	arr := kidsWithCandies(input,extra)
	if ! reflect.DeepEqual(arr,out) {
		t.Errorf("incorrect answer , got :%v , expected :%v",arr,out)
	}
}

func TestCandies1(t *testing.T) {
	input := []int {4,2,1,1,2}
	extra := 1
	out := [] bool{true,false,false,false,false}

	arr := kidsWithCandies(input,extra)
	if ! reflect.DeepEqual(arr,out) {
		t.Errorf("incorrect answer , got :%v , expected :%v",arr,out)
	}
}

func TestCandies2(t *testing.T) {
	input := []int {12,1,12}
	extra := 1
	out := [] bool{true,false,true}

	arr := kidsWithCandies(input,extra)
	if ! reflect.DeepEqual(arr,out) {
		t.Errorf("incorrect answer , got :%v , expected :%v",arr,out)
	}
}