package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

func main() {
	input := "hello"
	bytes, err := generate(input, 32.0, "SFNSMono.ttf")
	if err != nil {
		log.Fatalf("failed to genereate image for text %s with error %v", input, err)
	}

	// create image file
	imgFile, err := os.Create("output.png")
	if err != nil {
		log.Fatalf("failed to create image output file for text %s with error %v", input, err)
	}

	defer func() {
		err2 := imgFile.Close()
		if err2 != nil {
			panic(err2)
		}
	}()

	_, err = imgFile.Write(bytes)
	if err != nil {
		log.Fatalf("failed to write bytes to image for text %s with error %v", input, err)
	}

}

func generate(input string, fontSize float64, fontPath string) ([]byte, error) {

	// define forground and background color
	fgc := color.Black
	// bgc := color.White
	bgc := color.Transparent

	// Load font
	if &fontPath == nil {
		return nil, fmt.Errorf("fontPath is required")
	}

	ft, err := loadFont(fontPath)
	if err != nil {
		return nil, fmt.Errorf("failed to generate image due to error: %v", err)
	}

	fg := image.NewUniform(fgc)
	bg := image.NewUniform(bgc)

	rectangle := image.Rect(0, 0, 512, 64)
	canvas := image.NewRGBA(rectangle)
	draw.Draw(canvas, canvas.Bounds(), bg, image.Pt(0, 0), draw.Src)

	// raster text

	// split to array of lines if multiline text
	textLines := strings.Split(input, "\n")

	cxt := freetype.NewContext()
	cxt.SetDPI(72)
	cxt.SetFontSize(fontSize)
	cxt.SetFont(ft)
	cxt.SetDst(canvas)
	cxt.SetSrc(fg)
	cxt.SetClip(canvas.Rect)

	pt := freetype.Pt(32, 32)
	for _, line := range textLines {
		_, err := cxt.DrawString(line, pt)
		if err != nil {
			return nil, fmt.Errorf("failed to generate image due to error: %v", err)
		}
		pt.Y += cxt.PointToFixed(fontSize * 1.2)
	}

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, canvas); err != nil {
		return nil, fmt.Errorf("failed to generate image due to error: %v", err)
	}

	return buf.Bytes(), nil
}

// loadFont - return parsed truetyoe font for given path
// if fails return and error
func loadFont(path string) (*truetype.Font, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load font due to error: %v", err)
	}

	parsedFont, err := freetype.ParseFont(bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to load font due to error: %v", err)
	}
	return parsedFont, nil
}
