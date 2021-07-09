package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	greet(w, "Queiroz")
}

func main() {
	log.Fatal(http.ListenAndServe(":5000", http.HandleFunc(GreetHandler)))
}
