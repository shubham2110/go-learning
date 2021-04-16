package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}

	repeated = Repeat("a", 6)
	expected = "aaaaaa"

	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}

}

func BenchmarkRepeat(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	expected := Repeat("a", 5)
	fmt.Println(expected)
	// output: aaaaa
}
