package main

import (
	"slices"
	"testing"
)
func TestPrintMatrix(t *testing.T) {
	nums := [][]int{{1,2,3},{4,5,6},{7,8,9}} 
	got := PrintMatrix(nums)
	want := "1,2,3,4,5,6,7,8,9,"

	if(got!= want){
		t.Errorf("got : %v\n want: %v\n", got , want)
	}
}

func TestSumAll(t *testing.T) {
	t.Run("sum of 2 slices" , func(t *testing.T) {
		got := SumAll( []int{0,2,4} , []int{4,6,8} )
		want := []int{6,18}

		if !slices.Equal(got , want) {
			t.Errorf("got: %v\n want:%v ", got , want)
		}
	})

	t.Run("sum of any no of slices" , func(t *testing.T) {
		got := SumAll( []int{5,2,4} , []int{4,6,8} , []int{5,7,8})
		want := []int{11,18,20}

		if !slices.Equal(got , want) {
			t.Errorf("got: %v\n want:%v ", got , want)
		}
	})

	t.Run("passing empty slices" , func(t *testing.T) {
		got := SumAll([]int{}, []int{3, 4, 5})
		want := []int{0, 12}

		if !slices.Equal(got , want) {
			t.Errorf("got: %v\n want:%v ", got , want)
		}
	})
}