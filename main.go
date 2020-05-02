package main

import (
	"net/http"
	"log"
	"github.com/MieGorenk/s3uploader/handler"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/",handler.Home)
	router.HandleFunc("/resource", handler.PostFile).Methods("POST")
	
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT"},
    	AllowedHeaders: []string{"Content-Type"},
    	Debug:          true,
	})

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":3000",handler))
}
