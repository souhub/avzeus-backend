package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
	"github.com/souhub/avzeus-backend/pkg/route"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/actresses", route.Actresses)
	mux.HandleFunc("/api/wemen", route.Wemen)
	mux.HandleFunc("/api/recommendation", route.Recommendation)
	mux.HandleFunc("/api/result", route.Result)

	// CORSポリシー対策でフロントエンドからのAjax通信のみ許可する
	frontEndURL := os.Getenv("FRONTEND_URL")
	c := cors.New(cors.Options{
		AllowedOrigins: []string{frontEndURL},
	})
	// handler := cors.Default().Handler(mux)
	handler := c.Handler(mux)
	log.Fatalln(http.ListenAndServe(":8000", handler))
}
