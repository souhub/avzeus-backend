package main

import (
	"log"
	"net/http"

	"github.com/souhub/avzeus-backend/pkg/route"
)

func main() {
	http.HandleFunc("/actresses", route.Actresses)
	http.HandleFunc("/wemen", route.Wemen)
	http.HandleFunc("/selection", route.Selection)
	http.HandleFunc("/recommendation", route.Recommendation)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
