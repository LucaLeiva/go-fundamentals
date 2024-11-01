package arraysslices

import (
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
	got := SumAll([]int{1, 2}, []int{1, 9})
	want := []int{3, 10}

	// no se puede comparar los slices con != y ==
	// if !replect.DeepEqual(got, want) {
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	// asignar una funcion a una variable
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Suma de algunos slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("Suma slices vacios", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
