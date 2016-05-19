package main

import (
	"fmt"
	"net/http"
	"os"
)

func CreateHandleFn(name string) http.HandlerFunc {
	handle := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, this is the server with name: %s\n", name)
	}
	return http.HandlerFunc(handle)
}

func main() {
	http.ListenAndServe(":8080", CreateHandleFn(os.Args[1]))
}
