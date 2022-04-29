package main

import (
	"log"
	"net/http"
)

func brands(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id": "1", "name": "test brand 1"}, {"id": "1", "name": "test brand 2"}]`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"brand created"}`))
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{brand updated"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"brand deleted"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"not found"}`))
	}
}

func main() {
	http.HandleFunc("/api/v1/brands", brands)
	log.Fatal(http.ListenAndServe(":8010", nil))
}
