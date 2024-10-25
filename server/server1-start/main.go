package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func toJson(a interface{}) string {



}

func handleHome(w http.ResponseWriter, r *http.Request){


}



func main(){

	router := chi.NewRouter()
	router.Get("/", handleHome)
	router.Post("/", handleHome)

	http.ListenAndServe(":8000", nil)

}