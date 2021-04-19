package route

import (
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var sess = createSession()
var bucket = os.Getenv("S3_BUCKET_NAME")

func s3Upload(key string, file io.Reader) (err error) {
	bucket := os.Getenv("S3_BUCKET_NAME")
	uploader := s3manager.NewUploader(sess)
	uploadObject := s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	}
	_, err = uploader.Upload(&uploadObject)
	return err
}

func s3Delete(key string) (err error) {
	input := s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	svc := s3.New(sess)
	result, err := svc.DeleteObject(&input)
	if err != nil {
		return err
	}
	log.Println(result)
	return
}

func createSession() *session.Session {
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_DEFAULT_REGION")
	creds := credentials.NewStaticCredentials(accessKeyID, secretAccessKey, "")
	cfg := aws.Config{
		Credentials: creds,
		Region:      aws.String(region),
	}
	sess, err := session.NewSession(&cfg)
	if err != nil {
		log.Fatalln(err)
	}
	return sess
}
