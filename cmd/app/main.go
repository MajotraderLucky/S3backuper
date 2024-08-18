package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Config struct {
	AwsRegion    string `json:"awsRegion"`
	AwsAccessKey string `json:"awsAccessKey"`
	AwsSecretKey string `json:"awsSecretKey"`
	BucketName   string `json:"bucketName"`
	FilePath     string `json:"filePath"`
	KeyPath      string `json:"keyPath"`
	Endpoint     string `json:"endpoint"`
}

func loadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func main() {
	config, err := loadConfig("config/config.json")
	if err != nil {
		fmt.Println("Ошибка чтения конфигурации:", err)
		return
	}

	// Конфигурация доступа к AWS
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(config.AwsRegion),
		Credentials:      credentials.NewStaticCredentials(config.AwsAccessKey, config.AwsSecretKey, ""),
		Endpoint:         aws.String(config.Endpoint), // Установка кастомной конечной точки
		S3ForcePathStyle: aws.Bool(true),              // Важно для некоторых S3-совместимых провайдеров
	})
	if err != nil {
		fmt.Println("Ошибка создания сессии:", err)
		return
	}

	// Создание uploader
	uploader := s3manager.NewUploader(sess)

	// Открываем файл
	file, err := os.Open(config.FilePath)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	// Загрузка файла
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.BucketName),
		Key:    aws.String(config.KeyPath),
		Body:   file,
	})
	if err != nil {
		fmt.Println("Ошибка при загрузке файла:", err)
		return
	}
	fmt.Println("Файл успешно загружен")
}
