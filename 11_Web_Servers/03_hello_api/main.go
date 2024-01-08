package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tomasda7/hello_api/hello"
)

func main() {
	http.HandleFunc("/api/v1/hello", handleHello)

	const PORT string = ":7070"
	fmt.Println("Server started at <http://localhost" + PORT + ">")
	log.Fatal(http.ListenAndServe(PORT, nil))
}

var names = []string{}

type req struct {
	Name string `json:"name"`
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(names)
	}

	if r.Method == "POST" {
		var param req
		err := json.NewDecoder(r.Body).Decode(&param)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		greeting := hello.Greeting(param.Name, &names)
		w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(greeting)
	}
}
