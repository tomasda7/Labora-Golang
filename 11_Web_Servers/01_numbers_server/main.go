package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	numbers "github.com/tomasda7/numbers_server/numbers"
)

func main() {
	http.HandleFunc("/api/v1/sumEvens", handleSumEvens)
	http.HandleFunc("/api/v1/sumOdds", handleSumOdds)

	const PORT string = ":8080"
	fmt.Println("Server started at <http://localhost" + PORT + ">")
	log.Fatal(http.ListenAndServe(PORT, nil))
}

type res struct {
	Total int `json:"total"`
}

type req struct {
	Number int `json:"number"`
}


func handleSumEvens(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Method == "GET" {
		param := r.URL.Query().Get("number")
		num, err := strconv.Atoi(param)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var response res
		total := numbers.SumFirstEvens(num)
		response.Total = total

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}

	if r.Method == "POST" {
		var param req
		err := json.NewDecoder(r.Body).Decode(&param)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var response res
		total := numbers.SumFirstEvens(param.Number)
		response.Total = total

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func handleSumOdds(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Method == "GET" {
		param := r.URL.Query().Get("number")
		num, err := strconv.Atoi(param)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var response res
		total := numbers.SumFirstOdds(num)
		response.Total = total

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}

	if r.Method == "POST" {
		var param req
		err := json.NewDecoder(r.Body).Decode(&param)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var response res
		total := numbers.SumFirstOdds(param.Number)
		response.Total = total

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
