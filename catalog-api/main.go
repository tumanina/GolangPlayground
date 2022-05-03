package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func brandsEndpoints(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(getBrands())
	case "POST":
		w.WriteHeader(http.StatusCreated)
		reqBody, _ := ioutil.ReadAll(r.Body)
		var brand Brand
		json.Unmarshal(reqBody, &brand)
		addBrand(brand)
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		u, err := url.Parse(r.RequestURI)
		if err != nil {
			log.Fatal(err)
		}

		id, err := strconv.ParseInt(u.Query().Get("id"), 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		deleteBrand(int32(id))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"not found"}`))
	}
}

func main() {
	http.HandleFunc("/api/v1/brands", brandsEndpoints)
	log.Fatal(http.ListenAndServe(":8010", nil))
}
