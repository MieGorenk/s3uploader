package main

import (
	"net/http"
	"log"
	"github.com/MieGorenk/s3uploader/handler"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/",handler.Home)
	router.HandleFunc("/resource", handler.PostFile).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(headers, methods, origins)(router)))
}
