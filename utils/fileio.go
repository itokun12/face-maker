package utils

import (
	"io"
	"log"
	"mime/multipart"
	"os"
)

func SaveFile(file multipart.File, filename string, dir string) {
	out, err := os.Create("./images/" + dir + "/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
}
