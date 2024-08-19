package main

import (
	"fmt"
	"s3backuper/internal/config"
	"s3backuper/internal/services"
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

	err = services.UploadFile(uploader, cfg.FilePath, cfg.BucketName, cfg.KeyPath)
	if err != nil {
		fmt.Println("Ошибка при загрузке файла:", err)
		return
	}

	fmt.Println("Файл успешно загружен")
}
