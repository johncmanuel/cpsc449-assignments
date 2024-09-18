package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"strings"
)

type Person struct {
	Name   string `json:"name"`
	School string `json:"school"`
	OS     string `json:"os"`
	Editor string `json:"editor"`
	Lang   string `json:"lang"`
}

func main() {
	// 1st requirement
	http.HandleFunc("/hello/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		fmt.Fprintf(w, "Hello! My name is %q", html.EscapeString(r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]))
	})

	// 2nd requirement
	http.HandleFunc("/data/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}

		var person Person
		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}

		w.Header().Set("Content-Type", "application/json")

		res, err := json.Marshal(person)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		w.Write(res)
	})

	// 3rd requirement
	http.HandleFunc("/comb/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}

		var person Person
		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}

		// If method is GET, set editor to VSCode
		// Else, editor stays Neovim
		if r.Method == http.MethodGet {
			person.Editor = "VSCode"
		}

		w.Header().Set("Content-Type", "application/json")

		res, err := json.Marshal(person)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		w.Write(res)
	})

	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
