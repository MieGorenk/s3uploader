package handlers

import (
	"fmt"
	"net/http"

	"github.com/MieGorenk/s3uploader/helpers"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

)

func PostFile(w http.ResponseWriter, r *http.Request) {
	maxSize := int64(1024000)
	err := r.ParseMultipartForm(maxSize)
	if err != nil {
	   fmt.Fprintf(w, "Image too large. Max Size: %v", maxSize)
	   return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "Could not get uploaded file")
		return
	}
	defer file.Close()

	session, err := session.NewSession(&aws.Config{Region:aws.String("ap-southeast-1")})
	if err != nil {
		fmt.Fprintf(w, "Could not upload file")
	  }

	fileName, err := helpers.UploadFileToS3(session, file, fileHeader)
	if err != nil {
		fmt.Fprintf(w, "Could not upload file")
	}

	fmt.Fprintf(w, "Image uploaded successfully: %v", fileName)
}