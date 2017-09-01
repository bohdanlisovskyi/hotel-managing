package main

import (
	"log"
	"net/http"

	"os"

	"github.com/bohdanlisovskyi/hotel-managing/core/routes"
)

//Run REST API server
func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3001"
	}
	log.Fatal(http.ListenAndServe(":"+port, router.NewRouter()))
}
