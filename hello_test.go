package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkRandomInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}

func TestHello(t *testing.T) {

	t.Run("Test Success case", func(t *testing.T) {
		fmt.Println(Hello("Thiago"))
		// Output: Hello, Thiago.
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
