package main

import (
	"net/http"
	"log"
	"github.com/MieGorenk/s3uploader/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handlers.PostFile).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
