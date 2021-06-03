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
