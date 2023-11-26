package lib

import (
	models "backend/module"

	"github.com/h2non/bimg"
)

func FormatImage(image *[]byte, data models.ImageFormat) error {
	newImage, err := bimg.NewImage(*image).Convert(CheckDataFormat(data))
	if err != nil {
		return err
	}

	*image = newImage
	return nil
}

// function check dataFormat is valid and return bimg format type
func CheckDataFormat(data models.ImageFormat) bimg.ImageType {
	var format bimg.ImageType
	switch data.Format {
	case "jpeg":
		format = bimg.JPEG
	case "png":
		format = bimg.PNG
	case "webp":
		format = bimg.WEBP
	case "tiff":
		format = bimg.TIFF
	case "gif":
		format = bimg.GIF
	case "svg":
		format = bimg.SVG
	}
	return format
}
