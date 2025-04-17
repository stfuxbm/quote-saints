package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/stfuxbm/quote-saints/config"
	"github.com/stfuxbm/quote-saints/internal/database"
	"github.com/stfuxbm/quote-saints/internal/routes"
)

func main() {
	// Load environment variable dari file .env
	config.LoadEnv()

	// Koneksi ke MongoDB
	database.MongoConnect()

	// Setup semua route
	mux := routes.SetupRoutes()

	// Jalankan server di port 8080
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
