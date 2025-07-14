// package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	// Set your region and bucket
	region := "ap-south-1"
	bucket := "golang-project-bucket"
	key := "uploaded_file.txt"
	filePath := "output.txt"

	// Step 1: Load default config (from env, ~/.aws/, etc.)
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		fmt.Println("failed to load config:", err)
		return
	}

	// Step 2: Open the file to upload
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("failed to open file:", err)
		return
	}
	defer file.Close()

	// Step 3: Create S3 client
	client := s3.NewFromConfig(cfg)

	// Step 4: Upload file
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   file, // io.Reader
	})
	if err != nil {
		fmt.Println("failed to upload file:", err)
		return
	}

	fmt.Println("✅ File uploaded successfully!")
}
