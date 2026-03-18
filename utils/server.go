package utils

import (
	"encoding/json"
	"net/http"
)

func RunServer() {
	http.HandleFunc("/news", newsHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func newsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mp := map[string]string{
		"foo": "bar",
		"bar": "baz",
	}
	err := json.NewEncoder(w).Encode(mp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
