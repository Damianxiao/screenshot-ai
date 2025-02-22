package utils

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"math/rand"
	"os"
	"time"
)

func SaveImgToStatic(img *image.RGBA) (string, error) {
	wd, _ := os.Getwd()
	time := time.Now().Format("2006-01-02")
	path := wd + "/static/screenshot/" + time
	if _, err := os.Stat(path); err != nil {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			panic("create dir fail")
		}
	}
	fileName := fmt.Sprintf("/%s_%s_%d.png", "screenshot", time, rand.Int())
	filepath := path + fileName
	if _, err := os.Stat(filepath); err != nil {
		file, err := os.Create(filepath)
		if err != nil {
			return "", err
		}
		png.Encode(file, img)
	}

	return filepath, nil
}

func ImgToBase64(img string) (string, error) {
	imgBytes, err := os.ReadFile(img)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(imgBytes), nil
}
