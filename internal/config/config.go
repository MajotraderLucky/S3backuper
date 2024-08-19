package config

import (
	"encoding/json"
	"os"
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

func LoadConfig(path string) (*Config, error) {
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
