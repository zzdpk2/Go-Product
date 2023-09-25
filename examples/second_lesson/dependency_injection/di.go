package main

import (
	"fmt"
	"io"
	"net/http"
)

// func Greet(writer *bytes.Buffer, name string) {
func Greet(writer io.Writer, name string) {
	// fmt.Printf("Hello, %s", name)
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	// Greet(os.Stdout, "Elodie")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
	// buffer := bytes.Buffer{}
	// Greet(&buffer, "Chris")
}
