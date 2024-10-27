package arraysslices

import (
	"fmt"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Colleción de 5 numeros", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		// con [...] el compilador infiere el tamaño al crearlo
		// numbers := [...]int{1, 2, 3, 4, 5}

		got := SumArray(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("Collección de cualquier tamaño", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := SumSlice(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	fmt.Println(got)
	fmt.Println(want)

	// no se puede comparar los slices con != y ==
	// if !replect.DeepEqual(got, want) {
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
