package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	s3Client *s3.Client
	S3Bucket string
)

func init() {
	region := env("AWS_REGION", "us-east-1")
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			env("AWS_ACCESS_KEY_ID", "minioadmin"),
			env("AWS_SECRET_ACCESS_KEY", "minioadmin"),
			"",
		)),
	)
	if err != nil {
		panic(err)
	}
	opts := []func(*s3.Options){}
	if ep := os.Getenv("S3_ENDPOINT"); ep != "" {
		opts = append(opts, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(ep)
			o.UsePathStyle = true
		})
	}
	s3Client = s3.NewFromConfig(cfg, opts...)
	S3Bucket = env("S3_BUCKET", "goexpert-bucket-exemplo")
}

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()
}

func uploadFile(fileName string) {
	completeFileName := fmt.Sprintf("./tmp/%s", fileName)
	fmt.Println("Uploading file %s", completeFileName, "to bucket %s", S3Bucket)
	file, err := os.Open(completeFileName)
	if err != nil {
		fmt.Println("Error opening file %s", completeFileName)
		return
	}
	defer file.Close()
	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(S3Bucket),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		fmt.Println("Error uploading file %s", completeFileName)
		return
	}
	fmt.Println("File %s uploaded successfully", completeFileName)
}
