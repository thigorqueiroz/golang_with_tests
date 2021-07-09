package dependency

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestDi(t *testing.T) {
	buffer := bytes.Buffer{}
	name := "thiago"
	Greet(&buffer, name)

	got := buffer.String()
	want := "Hello, thiago \n"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

}

func TestGreetFromSo(t *testing.T) {
	Greet(os.Stdout, "Fullie")
}

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s \n", name)
}
