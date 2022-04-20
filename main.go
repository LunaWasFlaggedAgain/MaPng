package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func MapColorToRGBA(mapcolor byte) color.RGBA {
	basei, shadei := mapcolor/4, mapcolor%4
	base := BASECOLORS[basei]
	multiplier := SHADES[shadei]

	return color.RGBA{
		uint8((int(base.R) * multiplier) / 255),
		uint8((int(base.G) * multiplier) / 255),
		uint8((int(base.B) * multiplier) / 255),
		base.A}
}

func MapToImage(r io.Reader) (image.Image, error) {
	m, err := ParseMap(r)
	if err != nil {
		return nil, err
	}

	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{128, 128}})

	for i, block := range m.Colors {
		x := i % 128
		y := i / 128
		pix := MapColorToRGBA(block)

		img.Set(x, y, pix)
	}

	return img, nil
}

func main() {
	files, err := ioutil.ReadDir("./maps/")
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll("./images/", 0755)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println("processing", file.Name())

		cleanName := file.Name()[:strings.Index(file.Name(), ".")]

		file, err := os.Open("./maps/" + file.Name())
		if err != nil {
			fmt.Printf("got error while processing file %s: %v\n", file.Name(), err)
			continue
		}

		i, err := MapToImage(file)
		if err != nil {
			fmt.Printf("got error while processing file %s: %v\n", file.Name(), err)
			continue
		}

		file.Close()

		destFile, err := os.Create("./images/" + cleanName + ".png")
		if err != nil {
			fmt.Printf("got error while processing file %s: %v\n", file.Name(), err)
			continue
		}

		err = png.Encode(destFile, i)
		if err != nil {
			fmt.Printf("got error while processing file %s: %v\n", file.Name(), err)
			continue
		}

		destFile.Close()
	}
}
