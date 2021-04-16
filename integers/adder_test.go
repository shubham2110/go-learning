package integer

import "fmt"
import "testing"

func TestAdder(t *testing.T){

	got := Add(6,7)
	want := 13
	
	if got != want {
		t.Errorf("got %d want %d", got , want)
	}
}

func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Println(sum)
    // Output: 6
}