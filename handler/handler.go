// Handling all of the route regarding s3 operation
package handler

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/MieGorenk/s3uploader/helpers"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joho/godotenv"
)

// Construct for JSON messagess
type Response struct {
	URL string
}

type ErrorResponse struct {
	Error string
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func PostFile(w http.ResponseWriter, r *http.Request) {
	godotenv.Load()
	// Set maximum uploaded file to be 3 GB
	maxSize := int64(32221225472)
	err := r.ParseMultipartForm(maxSize)
	if err != nil {
	   fmt.Fprintf(w, err.Error())
	   return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "Could not get uploaded file")
		return
	}
	defer file.Close()

	// Create new session to AWS using saved credentials
	session, err := session.NewSession(&aws.Config{Region:aws.String("ap-southeast-1")})
	if err != nil {
		res := ErrorResponse{"Could not make connection to AWS"}
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(res)
		return
	  }

	// Uploading file to S3 using helper function
	// TODO add progress bar when uploading
	fileURL, err := helpers.UploadFileToS3(session, file, fileHeader)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Fprint(w, aerr.Code())
			return
		}
		// res := ErrorResponse{"No internet connection avalaible"}
		// w.WriteHeader(http.StatusConflict)
		// json.NewEncoder(w).Encode(res)
		// return
	}

	res := Response{fileURL}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
