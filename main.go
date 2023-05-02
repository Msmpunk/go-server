package main

import (
	"log"
	"net/http"

	"github.com/Msmpunk/go-server/routes"
)

func main() {
	r := routes.Router()
	log.Println("Starting server.")
	log.Fatal(http.ListenAndServe(":8081", r))
}
