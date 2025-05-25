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
		fmt.Println("Inside the get")
	}
	if r.Method == http.MethodPost {

		var x data
		err := json.NewDecoder(r.Body).Decode(&x)
		if err != nil {
			fmt.Println("Error tracing the request")
		} else {
			w.WriteHeader(http.StatusOK)
		}

		fmt.Println("The value of the id is ", x.Id, " and the value of the name is ", x.Name)
	}
}

func main() {
	http.HandleFunc("/", handleHome)
	http.ListenAndServe(":8080", nil)
}
