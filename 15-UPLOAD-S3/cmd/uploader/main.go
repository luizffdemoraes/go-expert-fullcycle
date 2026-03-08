package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

var (
	s3Client  *s3.Client
	S3Bucket  string
	uploadDir string
	wg        sync.WaitGroup
)

func init() {
	// Procura .env na pasta atual e nas pastas acima (para Code Runner / qualquer cwd)
	dir, _ := os.Getwd()
	for {
		if err := godotenv.Load(filepath.Join(dir, ".env")); err == nil {
			break
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
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
	opts := []func(*s3.Options){
		func(o *s3.Options) {
			o.BaseEndpoint = aws.String(env("S3_ENDPOINT", "http://localhost:9000"))
			o.UsePathStyle = true
		},
	}
	s3Client = s3.NewFromConfig(cfg, opts...)
	S3Bucket = env("S3_BUCKET", "goexpert-bucket-exemplo")
	uploadDir = filepath.Join(projectRoot(), env("UPLOAD_DIR", "tmp"))
}

func projectRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		return "."
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return dir
		}
		dir = parent
	}
}

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func main() {
	d, err := os.Open(uploadDir)
	if err != nil {
		panic(fmt.Sprintf("abrir diretório de upload %q: %v (rode o generator antes ou crie a pasta tmp)", uploadDir, err))
	}
	defer d.Close()

	// Struct é a menor unidade de memória que podemos usar para sincronização
	uploadControl := make(chan struct{}, 100)
	uploadControl <- struct{}{}
	for {
		files, err := d.Readdir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %v\n", err)
			continue
		}
		wg.Add(1)
		// Aguarda a liberação de um slot para subir o arquivo
		uploadControl <- struct{}{}
		go uploadFile(files[0].Name(), uploadControl)
	}
}

func uploadFile(fileName string, uploadControl <-chan struct{}) {
	defer wg.Done()
	completeFileName := filepath.Join(uploadDir, fileName)
	fmt.Printf("Uploading file %s to bucket %s\n", completeFileName, S3Bucket)
	file, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", completeFileName, err)
		<-uploadControl // Libera o slot para outro upload
		return
	}
	defer file.Close()
	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(S3Bucket),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		fmt.Printf("Error uploading file %s\n", completeFileName)
		fmt.Printf("  causa: %v\n", err)
		<-uploadControl // Libera o slot para outro upload
		return
	}
	fmt.Printf("File %s uploaded successfully\n", completeFileName)
	<-uploadControl // Libera o slot para outro upload
}
