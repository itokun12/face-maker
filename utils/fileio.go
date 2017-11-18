package utils

import (
	"io"
	"log"
	"mime/multipart"
	"os"
)

func SaveFile(file multipart.File, header multipart.FileHeader) {
	filename := header.Filename
	out, err := os.Create("./images/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
}
