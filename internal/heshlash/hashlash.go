package heshlash

import (
	"encoding/json"
	"fmt"
	"net/http"
	"homework_45/internal/models"
)

func GetExpectedHash(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var downloads []models.GoDownload
	err = json.NewDecoder(resp.Body).Decode(&downloads)
	if err != nil {
		return "", err
	}

	for _, download := range downloads {
		if download.Version == "go1.22.3" {
			for _, file := range download.Files {
				if file.Filename == "go1.22.3.src.tar.gz" {
					return file.Sha256, nil
				}
			}
		}
	}
	return "", fmt.Errorf("expected hash not found for go1.22.3.src.tar.gz")
}
