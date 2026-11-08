package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(map[string]string{
			"status": "ok", "service": "circula-api",
		}); err != nil {
			log.Printf("encode health response: %v", err)
		}
	})
	server := &http.Server{
		Addr: "127.0.0.1:8080", Handler: mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Print("Circula API: http://127.0.0.1:8080")
	log.Fatal(server.ListenAndServe())
}
