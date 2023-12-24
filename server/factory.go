package server

// Img Factory

import (
	"encoding/base64"
	"fmt"
	"image"
	_ "image/gif"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const SVG_TEMPLATE = `<?xml version="1.0" encoding="UTF-8"?>
<svg width="%d" height="%d" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" style="image-rendering: pixelated;">
    <title>Moon-Counter</title><g>%s</g>
</svg>`

const NUMBER_TEMPLATE = `<image x="%d" y="0" width="%d" height="%d" xlink:href="data:image/%s;base64,%s" />`

type Image struct {
	data   *string
	width  int
	height int
}

var images []Image

func LoadImgage(path string) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		log.Fatal("Cannot access:", path, err)
	}
	imageCfg, _, err := image.DecodeConfig(f)
	if err != nil {
		log.Fatal("Cannot decode image:", path, err)
	}
	bytes, _ := os.ReadFile(path)
	base64Str := base64.StdEncoding.EncodeToString(bytes)
	image := Image{
		data:   &base64Str,
		width:  imageCfg.Width,
		height: imageCfg.Height,
	}
	images = append(images, image)
}

func LoadAssets(path string) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			LoadImgage(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal("Cannot load assets:", err)
	}
}

func BuildCounterImg(c string) string {
	iTimes := 6 - len(c)
	for i := 0; i < iTimes; i++ {
		c = "0" + c
	}
	var numberTempletes string
	// Todo: Handle a situation if each image's dimentions are different
	for i, sDigit := range c {
		digit := strings.Index("0123456789", string(sDigit))
		img := images[digit]
		// Todo: Add more image type support, exclude .gif
		numberTempletes += fmt.Sprintf(NUMBER_TEMPLATE, i*img.width, img.width, img.height, "gif", *img.data)
	}
	return fmt.Sprintf(SVG_TEMPLATE, images[0].width*len(c), images[0].height, numberTempletes)
}
