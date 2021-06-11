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
			t.Errorf("GOT: %s is diff WANT: %s", got, want)
		}
	}
	t.Run("Test With Empty State", func(t *testing.T) {
		got := Hello("  ", "")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Test With Blank", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Elodie", "spanish")
		want := "Holla, Elodie."

		assertCorrectMessage(t, got, want)
	})

	t.Run("in english", func(t *testing.T) {
		got := Hello("Jonh", "english")
		want := "Hello, Jonh."

		assertCorrectMessage(t, got, want)
	})

	t.Run("in portuguese", func(t *testing.T) {
		got := Hello("Jonh", "portuguese")
		want := "Oi, Jonh."

		assertCorrectMessage(t, got, want)
	})
}
