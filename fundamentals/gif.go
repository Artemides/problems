package fundamentals

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"time"
)

var palette = []color.Color{color.White, color.RGBA{0xA8, 0xDF, 0x8E, 0x99}, color.RGBA{0xF3, 0xFD, 0xE8, 0x99}, color.RGBA{0xFF, 0xE5, 0xE5, 0x99}, color.RGBA{0xFF, 0xBF, 0xBF, 0x99}}

// const (
// 	whiteIndex = 0
// 	blackIndex = 1
// )

func Lissajous(out io.Writer, cicle float64) {
	const (
		res     = 0.00001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cicle*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t + freq + phase)
			rand.Seed(time.Now().UnixNano())
			colorIndex := uint8(rand.Intn(len(palette)))
			img.SetColorIndex(size+int(x*size+1), size+int(y*size+1), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
