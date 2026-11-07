package main
import("testing"
	
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
