package main

import (
	"fmt"
	"s3backuper/internal/workflow"
)

func main() {
	if err := workflow.RunUploadProcess("config/config.json"); err != nil {
		fmt.Println("Error:", err)
	}
}
