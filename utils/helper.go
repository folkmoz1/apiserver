package utils

import (
	"crypto/rand"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
)

func GenRandomBytes(size int) (blk []byte, err error) {
	blk = make([]byte, size)
	_, err = rand.Read(blk)
	return
}

func Loader(image multipart.File, imgExt string) (image.Image, error)  {
	switch imgExt {
	case "png":
		return png.Decode(image)
	default:
		return jpeg.Decode(image)
	}

}