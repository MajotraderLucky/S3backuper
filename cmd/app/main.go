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
