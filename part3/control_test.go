package main
import "testing"

func TestControl(t *testing.T) {
	t.Run("simple printing elements" , func (t * testing.T) {
		got := PrintSerial(5)
		want := "1,2,3,4,5,"

		if(got != want) {
			t.Errorf("\ngot : %v\nwant : %v\nincorrect output" , got , want)
		}
	})
}