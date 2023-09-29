package plot

import (
	"fmt"
	"image"
	"image/color"
	"math/cmplx"
	"sync"
	"time"
)

func Render() {
	now := time.Now()
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewNRGBA(image.Rect(0, 0, 1024, 1024))

	N := 64
	size := height / N
	done := make(chan struct{})
	var wg sync.WaitGroup

	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for py := size * i; py < size*(i+1); py++ {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)

					img.Set(px, py, mandelBrot(z))
				}
			}
			done <- struct{}{}
		}(i)
	}
	go func() {
		wg.Wait()
		close(done)
	}()

	for range done {

	}
	// png.Encode(os.Stdout, img)
	fmt.Println(time.Since(now))
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
			return color.RGBA{r, g, b, 255}
		}

	}
	return color.RGBA{0, 0, 0, 255}
}
