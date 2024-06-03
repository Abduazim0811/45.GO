package main

import (
	"fmt"
	"homework_45/internal/heshlash"
)

func main() {
	goURL := "https://go.dev/dl/go1.22.3.src.tar.gz"
	archivePath := "go1.22.3.src.tar.gz"
	err := heshlash.DownloadFile(goURL, archivePath)
	if err != nil {
		fmt.Printf("Error downloading file: %v\n", err)
		return
	}

	jsonURL := "https://go.dev/dl/?mode=json"
	expectedHash, err := heshlash.GetExpectedHash(jsonURL)
	if err != nil {
		fmt.Printf("Error getting expected hash: %v\n", err)
		return
	}

	calculatedHash, err := heshlash.CalculateSHA256(archivePath)
	if err != nil {
		fmt.Printf("Error calculating hash: %v\n", err)
		return
	}

	if calculatedHash == expectedHash {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
