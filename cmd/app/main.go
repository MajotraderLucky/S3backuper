package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	// Конфигурация доступа к AWS
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ru-1"), // Укажите ваш регион
		Credentials: credentials.NewStaticCredentials("YOUR_ACCESS_KEY", "YOUR_SECRET_KEY", ""),
	})
	if err != nil {
		fmt.Println("Ошибка создания сессии:", err)
		return
	}

	// Создание uploader
	uploader := s3manager.NewUploader(sess)

	// Открываем файл
	file, err := os.Open("your-archive.tar.gz")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	// Загрузка файла
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("your-bucket-name"),
		Key:    aws.String("path/to/your-archive.tar.gz"),
		Body:   file,
	})
	if err != nil {
		fmt.Println("Ошибка при загрузке файла:", err)
		return
	}
	fmt.Println("Файл успешно загружен")
}
