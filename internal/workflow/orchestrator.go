package workflow

import (
	"fmt"
	"s3backuper/internal/config"
	"s3backuper/internal/services"
)

// RunUploadProcess coordinates the entire upload process
func RunUploadProcess(configPath string) error {
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
