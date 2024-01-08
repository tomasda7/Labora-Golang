package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/tomasda7/counter_server/counter"
)

func main() {
	http.HandleFunc("/api/v1/inc", handleInc)
	http.HandleFunc("/api/v1/dec", handleDec)
	http.HandleFunc("/api/v1/curr", handleCurr)

	const PORT string = ":9090"
	fmt.Println("Server started at <http://localhost" + PORT + ">")
	log.Fatal(http.ListenAndServe(PORT, nil))
}

var count int = 0

type res struct {
	Current int `json:"current"`
}

func handleInc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var inc sync.Mutex

	inc.Lock()
	count = counter.Increment(count)
	inc.Unlock()

	w.WriteHeader(http.StatusOK)
}

func handleDec(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var dec sync.Mutex
	dec.Lock()
	count = counter.Decrement(count)
	dec.Unlock()

	w.WriteHeader(http.StatusOK)
}

func handleCurr(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var response res
	response.Current = count

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
