package main
import "testing"

func TestHello(t *testing.T ) {
	t.Run("no name is given " ,  func (t *testing.T) {
		got := hello("")
		want := "Hello , World"
		assertCorrectMessage(t,got,want)
	})

	t.Run("a name is given " ,  func (t *testing.T) {
		got := hello("Niggers")
		want := "Hello , Niggers"
		assertCorrectMessage(t,got,want)
	})
	
}

func assertCorrectMessage(t *testing.T , got , want string){
	t.Helper()
	if(got != want){
		t.Errorf("\n got this : %v \n want this : %v \n" , got,want)
	}
}