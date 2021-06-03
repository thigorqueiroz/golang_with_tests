package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("Test Success case", func(t *testing.T) {
		got := Hello("Thiago")
		want := "Hello, Thiago."

		if got != want {
			t.Errorf("Got: %q is not equal Want:  %q", got, want)
		}
	})

	t.Run("Test With Empty State", func(t *testing.T) {
		got := Hello("  ")
		want := "Hello World"

		if got != want {
			t.Errorf("'Got' %q is not equal 'Want' %q ", got, want)
		}
	})

	t.Run("Test With Blank", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"

		if got != want {
			t.Error("There is a diference between GOT and WANT")
		}
	})
}
