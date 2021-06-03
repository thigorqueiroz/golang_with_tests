package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Thiago")

	want := "Hello, Thiago."

	if got != want {
		t.Errorf("Got: %q is not equal Want:  %q", got, want)
	}

	fmt.Printf("Got: %q is EQUAL want: %q ", got, want)
}

func TestHelloWithEmptyState(t *testing.T) {
	got := Hello("  ")
	want := "Hello World"

	if got != want {
		t.Errorf("'Got' %q is not equal 'Want' %q ", got, want)
	}

	fmt.Println("'Got' is equal 'Want'")
}

func TestHelloWithBlank(t *testing.T) {
	got := Hello("")
	want := "Hello World"

	if got != want {
		t.Error("There is a diference between GOT and WANT")
	}
}
