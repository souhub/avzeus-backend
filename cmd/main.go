package main

import (
	"log"
	"net/http"

	"github.com/souhub/avzeus-backend/pkg/route"
)

func main() {
	http.HandleFunc("/actresses", route.Actresses)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
