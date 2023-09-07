package plot

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func Render() {

	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewNRGBA(image.Rect(0, 0, 1024, 1024))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelBrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelBrot(z complex128) color.RGBA {
	const iterations = 128
	var v complex128
	for i := uint8(0); i < iterations; i++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {

			r := uint8(255 - real(v))
			g := uint8(255 - i)
			b := uint8(255 - i)
			fmt.Println(r, g, b)
			return color.RGBA{r, g, b, 255}
		}

	}
	return color.RGBA{0, 0, 0, 255}
}
