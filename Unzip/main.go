package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	zipReader, err := zip.OpenReader("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer zipReader.Close()

	for _, file := range zipReader.File {
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()

		targedDir := "./"
		extractedFilePath := filepath.Join(targedDir, file.Name)

		if file.FileInfo().IsDir() {
			log.Println("Creating directory: ", extractedFilePath)
			err := os.MkdirAll(extractedFilePath, file.Mode())
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Println("Extracting file:", file.Name)
			outputFile, err := os.OpenFile(extractedFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, file.Mode())
			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()

			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				log.Fatal(err)
			}
		}

	}
}
