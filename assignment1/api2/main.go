package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name   string `json:"name"`
	School string `json:"school"`
	OS     string `json:"os"`
	Editor string `json:"editor"`
	Lang   string `json:"lang"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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

	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
