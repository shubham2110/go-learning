package structs

import "testing"

func TestPerimeter(t *testing.T) {

	assert := func(t testing.TB, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %.2f   want %.2f ", got, want)
		}
	}

	t.Run("Testing got 10 and 10 ", func(t *testing.T) {
		got := Perimeter(Rectangle{10.0, 10.0})
		want := 40.0
		assert(t, got, want)

	})

	t.Run("Testing got 0.125 and 0.125  ", func(t *testing.T) {
		rectangle := Rectangle{0.125, 0.125}
		got := Perimeter(rectangle)
		want := 0.50
		assert(t, got, want)

	})

}

func TestArea(t *testing.T) {

	assert := func(t testing.TB, got, want float64) {
		if got != want {
			//t.Errorf("got %.2f  want %.2f", got, want)
			t.Errorf("got %g  want %g", got, want)
		}
	}

	t.Run("Checking for 10.0 and 20.0 for rectangle ", func(t *testing.T) {

		got := Area(Rectangle{10.0, 20.0})
		want := 200.00
		assert(t, got, want)
	})

	// t.Run("Checking for Circle with radius 10 ", func(t *testing.T) {

	// 	got := Area(Circle{10.0})
	// 	want := 314.1592653589793
	// 	assert(t, got, want)
	// })

}
