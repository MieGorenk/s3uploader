package main

import (
	"net/http"
	"log"
	"github.com/MieGorenk/s3uploader/handlers"
	"github.com/gorilla/mux"
)

const (
	AWS_S3_REGION = "ap-southeast-1"
	AWS_S3_BUCKET = "ezza-videos"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handlers.PostFile)
	log.Fatal(http.ListenAndServe(":8080", router))
}
