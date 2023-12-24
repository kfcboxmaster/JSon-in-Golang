package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", JSONRequest)
	log.Print("starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func JSONRequest(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var body map[string]any
	resp := make(map[string]string)

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}
	val, check := body["message"]
	if check {
		fmt.Println(val)
		resp["status"] = "success"
		resp["message"] = "Data successfully received"
	} else {
		resp["status"] = "404"
		resp["message"] = "Invalid JSON message"
	}

	w.Header().Set("Content-Type", "json")
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

/*
first attempt of assignment. query were used.
*/

func JSONRequestFirstAttempt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	queryParams := r.URL.Query()
	if queryParams.Has("message") {
		jsonResponseSuccess := []byte(`
	{
		"status": "success",
		"message": "Data successfully received"
	}
`)
		w.Write(jsonResponseSuccess)
	} else {
		jsonResponseError := []byte(`
	{
		"status": "400",
		"message": "Invalid JSON message"
	}
`)
		w.Write(jsonResponseError)
	}
}
