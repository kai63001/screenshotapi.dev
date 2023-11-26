package lib

import "github.com/h2non/bimg"

func FormatImage(image []byte) ([]byte, error) {
	newImage, err := bimg.NewImage(image).Convert(bimg.PNG)
	if err != nil {
		return image, err
	}
	return newImage, nil
}
