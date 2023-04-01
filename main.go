package main

import (
	"fmt"
	"net/http"
)

func handelerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, 这里是goblog</h1>")
}

func main() {
	http.HandleFunc("/", handelerFunc)
	http.ListenAndServe(":3000", nil)
}
