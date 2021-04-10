package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/souhub/avzeus-backend/pkg/route"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/actresses", route.Actresses).Methods("GET")
	r.HandleFunc("/wemen", route.Wemen).Methods("GET")
	r.HandleFunc("/recommendation", route.GetRecommendationActresses).Methods("POST")
	r.HandleFunc("/result", route.Result).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
