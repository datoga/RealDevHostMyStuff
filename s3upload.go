package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	S3_REGION        = "us-east-1"
	S3_BUCKET        = "realdevdatoga"
	S3_BUCKET_PREFIX = "files"
)

func uploadToS3(path string) {
	// Create a single AWS session (we can re use this if we're uploading many files)
	s, err := session.NewSession(&aws.Config{Region: aws.String(S3_REGION)})

	if err != nil {
		panic(err)
	}

	if isDirectory(path) {
		uploadDirToS3(s, path)
	} else {
		uploadFileToS3(s, path)
	}
}

func uploadDirToS3(sess *session.Session, dirPath string) {
	fileList := getFileList(dirPath)

	for _, file := range fileList {
		uploadFileToS3(sess, file)
	}
}

func uploadFileToS3(sess *session.Session, filePath string) {
	if debug {
		fmt.Println("upload " + filePath + " to S3")
	}
	// An s3 service
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Failed to open file", file, err)
		panic(err)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.

	fileDirectory, _ := filepath.Abs(filePath)
	key := S3_BUCKET_PREFIX + fileDirectory

	_, err = s3.New(sess).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(S3_BUCKET),
		Key:                  aws.String(key),
		ACL:                  aws.String("public-read"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	if err != nil {
		fmt.Printf("Failed to upload data to %s/%s, %s\n",
			S3_BUCKET, filePath, err.Error())
		return
	}

	urlF := "https://%s.s3.amazonaws.com/%s"
	url := fmt.Sprintf(urlF, S3_BUCKET, key)

	fmt.Println(url)
}
