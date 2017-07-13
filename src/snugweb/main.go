package main

import (
	"log"
	"net/http"
	"snugweb/handlers"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	hello := handlers.NewHello()

	router.HandleFunc("/", hello.ServeHTTP)
	router.HandleFunc("/data", handlers.DataHandler)

	log.Fatal(http.ListenAndServe(":9000", router))
}
