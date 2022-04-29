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
		w.Write([]byte(`[{"id": "1", "name": "test user 1"}, {"id": "1", "name": "test user 2"}]`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"user created"}`))
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{user updated"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"user deleted"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"not found"}`))
	}
}

func main() {
	http.HandleFunc("/api/v1/users", brands)
	log.Fatal(http.ListenAndServe(":8020", nil))
}
