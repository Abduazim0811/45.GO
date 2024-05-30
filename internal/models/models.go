package models

type GoDownload struct {
	Version string `json:"version"`
	Files   []struct {
		Filename string `json:"filename"`
		Sha256   string `json:"sha256"`
	} `json:"files"`
}
