package main

import (
	"Tracky/config"
	"Tracky/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	port := getPort()
	db := config.InitDb(":memory:")
	defer db.Close()

	router := routes.RegisterRoutes()

	log.Printf("Server is running on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func getPort() string {
	if port, exists := os.LookupEnv("PORT"); exists {
		return port
	}
	return "8080"
}
