// Helper function that will be used on handlers.go
package helpers

import (
	"mime/multipart"
	"bytes"
	"net/http"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joho/godotenv"

)

func UploadFileToS3(s *session.Session, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	godotenv.Load()

	size := fileHeader.Size
  	buffer := make([]byte, size)
	file.Read(buffer)
	  
	// config settings: this is where you choose the bucket,
	// filename, content-type and storage class of the file
	// you're uploading
	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String("ezza-videos"),
		Key:                  aws.String(fileHeader.Filename),
		ACL:                  aws.String("public-read"),// could be private if you want it to be access by only authorized users
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(int64(size)),
		ContentType:        aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("STANDARD"),
	})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3-%s.amazonaws.com/%s", "ezza-videos", "ap-southeast-1", fileHeader.Filename)

 	return url, err
}
