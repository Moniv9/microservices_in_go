package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/driver", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":3000", nil)
}
