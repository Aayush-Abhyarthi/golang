package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type data struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")
		fmt.Println("The id value is ", id)
	}
	if r.Method == http.MethodPost {
		var x data
		json.NewDecoder(r.Body).Decode(&x)
		fmt.Println("The id value is ", x.Id, " and the name value is ", x.Name)
	}
}

func main() {
	http.HandleFunc("/", handleHome)
	http.ListenAndServe(":8080", nil)
}
