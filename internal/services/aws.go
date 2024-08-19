package services

import (
	"fmt"
	"s3backuper/internal/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func InitS3Uploader(cfg *config.Config) (*s3manager.Uploader, error) {
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
