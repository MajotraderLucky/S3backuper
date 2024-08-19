package main

import (
	"fmt"
	"os"
	"s3backuper/internal/config"
	"s3backuper/internal/services"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		fmt.Println("Ошибка чтения конфигурации:", err)
		return
	}

	uploader, err := services.InitS3Uploader(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = uploadFile(uploader, cfg.FilePath, cfg.BucketName, cfg.KeyPath)
	if err != nil {
		fmt.Println("Ошибка при загрузке файла:", err)
		return
	}

	fmt.Println("Файл успешно загружен")
}

func uploadFile(uploader *s3manager.Uploader, filePath, bucketName, keyPath string) error {
	// Открываем файл
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Загрузка файла
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(keyPath),
		Body:   file,
	})
	if err != nil {
		return err
	}

	return nil
}
