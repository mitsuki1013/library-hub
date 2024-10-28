package main

import (
	"encoding/json"
	"fmt"
	"io"
	"library-hub/app/adapter/api/schema"
	"net/http"
)

func main() {
	http.HandleFunc("/company", CreateCompany)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var requestData schema.Company
	if err = json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Error parsing request json", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(requestData); err != nil {
		http.Error(w, "Error encoding response json", http.StatusInternalServerError)
		return
	}
}
