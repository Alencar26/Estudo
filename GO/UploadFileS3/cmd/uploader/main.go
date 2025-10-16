package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Cliet  *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
)

// ESSA FUNC É EXECUTADA ANTES DE MAIN POR DEFAULT NO GO
func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Endpoint: aws.String("http://localhost:4566"),
			Region:   aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(
				"localstack",
				"localstack",
				"",
			),
			S3ForcePathStyle: aws.Bool(true), //linha adicionada para não falha o localstack.
		},
	)
	if err != nil {
		panic(err)
	}
	s3Cliet = s3.New(sess)
	s3Bucket = "meu-bucket"
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	//channel para controle de quantas goroutines vou ter em paralelo
	uploadControl := make(chan struct{}, 100) //strutct é minha menor unidade por isso o uso dela
	errorFileUpload := make(chan string, 10)

	go func() {
		for {
			select {
			case filename := <-errorFileUpload:
				uploadControl <- struct{}{}
				wg.Add(1)
				go uploadFile(filename, uploadControl, errorFileUpload)
			}
		}
	}()

	for {
		files, err := dir.ReadDir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %s\n", err)
			continue
		}
		wg.Add(1)
		uploadControl <- struct{}{} //recebe um struct vazia a cada iteração
		go uploadFile(files[0].Name(), uploadControl, errorFileUpload)
	}
	wg.Wait()
}

func uploadFile(filename string, uploadControl <-chan struct{}, errorFileUpload chan<- string) {
	defer wg.Done()
	completeFileName := fmt.Sprintf("./tmp/%s", filename)
	fmt.Printf("Uploading file %s to bucket %s\n", completeFileName, s3Bucket)

	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Error opening file %s\n", completeFileName)
		<-uploadControl //esvaziando channel em caso de erro
		errorFileUpload <- filename
		return
	}
	defer f.Close()

	_, err = s3Cliet.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(s3Bucket),
		Key:         aws.String(filename),
		Body:        f,
		ContentType: aws.String("text"),
	})
	if err != nil {
		fmt.Printf("Error uploading file %s\n", completeFileName)
		<-uploadControl //esvaziando channel em caso de erro
		errorFileUpload <- filename
		return
	}

	fmt.Printf("File %s iploaded successfully\n", completeFileName)
	<-uploadControl //esvaziando channel em caso de erro
}
