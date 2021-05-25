package structs

import "testing"

func TestPerimeter(t *testing.T) {

	testcases := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle testing", shape: Rectangle{Width: 12, Height: 6}, want: 36.00},
		{name: "circle testing", shape: Circle{Radius: 10}, want: 62.83185307179586},
	}

	assert := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %g   want %g ", got, want)
		}
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			assert(t, testcase.shape, testcase.want)
		})
	}

}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle testing", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "circle testing", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "triangle testing", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	assert := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g  want %g", got, want)
		}
	}

	for _, testcase := range areaTests {
		t.Run(testcase.name, func(t *testing.T) {
			assert(t, testcase.shape, testcase.want)
		})
	}

}
