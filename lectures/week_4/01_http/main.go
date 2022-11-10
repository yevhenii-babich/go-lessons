package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/gopherize", gopherizeMe)
	http.HandleFunc("/greeting/new", textOut)
	http.HandleFunc("/greeting", textOut)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Printf("http server error: %v", err)
	}
}

func textOut(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte("Hello," + name)); err != nil {
		log.Printf("can't write: %v", err)
	}
}

func gopherizeMe(w http.ResponseWriter, r *http.Request) {
	body := &struct {
		Name     string `json:"name"`
		FullName string `json:"full_name"`
	}{}
	log.Printf("%+v", *r)
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	case http.MethodOptions:
		w.Header().Add("access-control-allow-methods", "POST,PUT")
		w.Header().Add("access-control-allow-origin", "*")
		w.WriteHeader(http.StatusOK)
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Printf("error decode: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body.FullName = body.Name + " Gopher"

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	raw, err := json.MarshalIndent(&body, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(raw)
	return
}
