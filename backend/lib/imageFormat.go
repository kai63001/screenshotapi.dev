package lib

import (
	"github.com/h2non/bimg"
)

func FormatImage(image *[]byte, formatImage string) error {
	newImage, err := bimg.NewImage(*image).Convert(CheckDataFormat(formatImage))
	if err != nil {
		return err
	}

	*image = newImage
	return nil
}

func ImageQuality(image *[]byte, quality int) error {
	newImage, err := bimg.NewImage(*image).Process(bimg.Options{
		Quality: quality,
	})
	if err != nil {
		return err
	}

	*image = newImage
	return nil
}

// function check dataFormat is valid and return bimg format type
func CheckDataFormat(formatImage string) bimg.ImageType {
	var format bimg.ImageType
	switch formatImage {
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
