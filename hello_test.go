package main

import "testing"

/*
	Escribir un test
	Hacer que compile
	Correr el test y que falle con un mensaje significativo
	Escribir suficiente codigo para que el test pase
	Refactorizar
*/

func TestHello(t *testing.T) {
	t.Run("Decirle hola a la gente", func(t *testing.T) {
		got := Hello("Luca")
		want := "Hola, Luca"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Deci 'Hola, Mundo' cuando se recibe un string vacio", func(t *testing.T) {
		got := Hello("")
		want := "Hola, Mundo"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// con esto le digo a los tests que esta es una funcion auxiliar, asi cuando hay un error
	// me lo indica en la linea del test que fallo
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
