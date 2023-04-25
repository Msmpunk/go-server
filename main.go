package main

import (
	"github.com/Msmpunk/go-server/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.Router()
	log.Println("Starting server.")
	log.Fatal(http.ListenAndServe(":8080", r))
}
