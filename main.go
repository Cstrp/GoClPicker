package main

import (
	"bufio"
	"fmt"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"sort"

	"github.com/disintegration/imaging"
)

type ColorData struct {
	color string
	length int
}



func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter path to image: ")

	path, _ := reader.ReadString('\n')
	path = path[:len(path)-1]

	colors, err := findColors(path, 5)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, c := range colors {
		displayColors(c.color)
	}


}

func findColors(path string, length int) ([]ColorData, error) {
	img, err := imaging.Open(path)
	if err != nil {
		return nil, err
	}

	colorMap := make(map[string]int)
	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			px := img.At(x,y)
			r, g, b, _ := px.RGBA()
			rgb := color.RGBA{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8), A: 255}
			hex := convertToHex(rgb)
			colorMap[hex]++
		}
	}

	colorInfo := make([]ColorData, 0, len(colorMap))

	for cl, ct := range colorMap {
		colorInfo = append(colorInfo, ColorData{
			color: cl,
			length: ct,
		})
	}

	sort.Slice(colorInfo, func(i, j int) bool {
		return colorInfo[i].length > colorInfo[j].length
	})

	if length < len(colorInfo) {
		colorInfo = colorInfo[:length]
	}

	return colorInfo, nil
}

func displayColors(color string) {
	fmt.Println(color)
}

func convertToHex(color color.RGBA) string {
	return fmt.Sprintf("#%02X%02X%02X", color.R, color.G, color.B)
}
