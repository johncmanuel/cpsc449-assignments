package main

import (
	"fmt"
	"html"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		fmt.Fprintf(w, "Hello! My name is %q", html.EscapeString(r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]))
	})

	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
