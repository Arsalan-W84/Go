package main
import "testing"

func TestHello(t *testing.T ) {
	got := hello("Niggers")
	want := "Hello , Niggers"

	if(got != want){
		t.Errorf("\n got this : %v \n want this : %v \n" , got,want)
	}
}