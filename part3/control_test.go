package main

import (
	"testing"
)

func TestControl(t *testing.T) {
	t.Run("simple printing elements with for loop" , func (t * testing.T) {
		got := PrintSerial(5)
		want := "1,2,3,4,5"

		assertErrorMessage(t,got,want);
	})
	t.Run("simple printing elements with while style loop" , func (t * testing.T) {
		got := Printwhileloop(5)
		want := "1,2,3,4,5"

		assertErrorMessage(t,got,want);
	})
	t.Run("Printing all elements of an array" , func (t *testing.T) {
		str := []string{"Hi" , "Hello" , "Thank You" , "Please"}
		got := PrintArrayElements(str)
		want := "Hi,Hello,Thank You,Please"
		assertErrorMessage(t,got,want);
	})
}

func assertErrorMessage (t* testing.T , got , want string) {
	t.Helper()
	if(got != want) {
			t.Errorf("\ngot : %v\nwant : %v\nincorrect output" , got , want)
	}
}

var result string
func Benchmark1(b * testing.B) {
	str := []string{"a" , "svss" , "sfsf" , "Psdf" , "swsk" , "swsn", "swejo"}
	for b.Loop() {
		result = PrintArrayElements(str)
	}
}

func Benchmark2(b * testing.B) {
	str := []string{"a" , "svss" , "sfsf" , "Psdf" , "swsk" , "swsn", "swejo"}
	for b.Loop() {
		result = PrintArray_without_builder(str)
	}
}


//OBSERVATION : STRING BUILDER IS FASTER THAN REGULAR STRING ADDITION : 
//Reason : 
// Strings in Go are immutable, meaning every concatenation, such as in our Repeat function, 
// involves copying memory to accommodate the new string. This impacts performance, particularly during heavy string concatenation. 

//BENCHMARK SCORES : 
// Benchmark1-16           14989315                76.97 ns/op           56 B/op          3 allocs/op
// Benchmark2-16            4380668               273.1 ns/op           232 B/op         12 allocs/op
// PASS
