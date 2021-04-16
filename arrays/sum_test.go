package arrays

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("checking first test case", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 6, 7})
		want := []int{2, 13}

		checkSums(t, got, want)

	})

	t.Run("safely run with empty sliced", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 6, 7})
		want := []int{0, 13}
		checkSums(t, got, want)
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3, 4}, []int{1, 2, 3})
	want := []int{10, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v  want %v", got, want)
	}
}

func TestSum(t *testing.T) {

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}

	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}
