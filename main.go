package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/greet", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello World!")
	})

	http.ListenAndServe(":8000", nil)
}
