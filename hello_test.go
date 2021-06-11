package main

import (
	"math/rand"
	"testing"
)

func BenchmarkRandomInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	t.Run("Test With Empty State", func(t *testing.T) {
		got := Hello("  ")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Test With Blank", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})
}
