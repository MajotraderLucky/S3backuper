package main

import (
	"fmt"
	"s3backuper/internal/config"
	"s3backuper/internal/services"
)

func main() {
	if err := runUploadProcess("config/config.json"); err != nil {
		fmt.Println("Error:", err)
	}
}

func runUploadProcess(configPath string) error {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("failed to read configuration: %w", err)
	}

	uploader, err := services.InitS3Uploader(cfg)
	if err != nil {
		return fmt.Errorf("failed to initialize S3 uploader: %w", err)
	}

	err = services.UploadFile(uploader, cfg.FilePath, cfg.BucketName, cfg.KeyPath)
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}

	fmt.Println("File successfully uploaded")
	return nil
}
