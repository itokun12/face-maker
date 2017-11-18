package utils

import (
	"gopkg.in/gographics/imagick.v3/imagick"
	"os"
)

func Resize(filename string) {
	imagick.Initialize()
	mw := imagick.NewMagickWand()

	defer imagick.Terminate()
	var err error

	err = mw.ReadImage("/images/original/" + filename)
	if err != nil {
		panic(err)
	}

	width := mw.GetImageWidth()
	height := mw.GetImageHeight()

	// Calculate half the size
	hWidth := uint(width / 2)
	hHeight := uint(height / 2)

	err = mw.ResizeImage(hWidth, hHeight, imagick.FILTER_LANCZOS)
	if err != nil {
		panic(err)
	}

	// Set the compression quality to 95 (high quality = low compression)
	err = mw.SetImageCompressionQuality(95)
	if err != nil {
		panic(err)
	}

	err = mw.DisplayImage(os.Getenv("DISPLAY"))
	if err != nil {
		panic(err)
	}
}
