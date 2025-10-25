package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "my client")
}

func main() {
	http.ListenAndServe("localhost:5001", http.HandlerFunc(MyGreetHandler))
}
