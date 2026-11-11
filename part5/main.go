//-------- ARRAYS SLICES AND MAPS  IN GO------

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PrintMatrix(nums [][]int) string {
	var op strings.Builder
	
	for _ , row := range nums {
		for _, value := range row {
			op.WriteString( strconv.Itoa(value) )
			op.WriteString(",")
		}
	}
	return op.String()
}

//func to return the sum of all the numbers of all the slices passed 
//input : a varying number of slices
//output : slice containing sum of all the slices passed
func SumAll( nums ...[]int ) []int {
	var op []int
	var sum int
	for _, slice := range nums {
		sum = 0
		for _, val := range slice {
			sum+=val
		}
		op = append(op, sum)
	}
	return op
}
func main() {
	
//static decalaration
	var numbers [4]int
	numbers[0] = 1
	fmt.Println(numbers)
//dynamic declaration
	nums := [4]int{2,4,5,7}

	for i:=0;i<len(nums);i++{
		fmt.Printf("%v , " , nums[i])
	}
	fmt.Println("")
	
}