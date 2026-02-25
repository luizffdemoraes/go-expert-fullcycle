package main

import (
	"context"
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
}
