package main

import (
	"log"
	"net/http"

	"github.com/souhub/av-zeus/pkg/route"
)

func main() {
	http.HandleFunc("/", route.Actresses)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
