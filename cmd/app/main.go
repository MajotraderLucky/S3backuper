package main

import (
	"fmt"
	"os"
	"s3backuper/internal/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func initS3Uploader(cfg *config.Config) (*s3manager.Uploader, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(cfg.AwsRegion),
		Credentials:      credentials.NewStaticCredentials(cfg.AwsAccessKey, cfg.AwsSecretKey, ""),
		Endpoint:         aws.String(cfg.Endpoint),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %w", err)
	}

	uploader := s3manager.NewUploader(sess)
	return uploader, nil
}

func main() {
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		fmt.Println("Ошибка чтения конфигурации:", err)
		return
	}

	uploader, err := initS3Uploader(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Открываем файл
	file, err := os.Open(cfg.FilePath)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	// Загрузка файла
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(cfg.BucketName),
		Key:    aws.String(cfg.KeyPath),
		Body:   file,
	})
	if err != nil {
		fmt.Println("Ошибка при загрузке файла:", err)
		return
	}
	fmt.Println("Файл успешно загружен")
}
