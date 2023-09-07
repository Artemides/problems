package plot

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func Draw() {
	filePath := "plot.txt"
	tag := fmt.Sprintf("<svg' "+
		"style='stroke:grey; fill:white; stroke-width: 0.7'"+
		"width='%d' height='%d' >", width, height)
	writeToFile(filePath, []byte(tag), false)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			points := fmt.Sprintf("<polygon points='%g, %g %g, %g %g, %g %g, %g ' />\n", ax, ay, bx, by, cx, cy, dx, dy)
			writeToFile(filePath, []byte(points), false)
		}
	}
	writeToFile(filePath, []byte("</svg>"), true)
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func writeToFile(filePath string, content []byte, close bool) bool {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error Reading file...", err)
		return false
	}
	if close {
		defer file.Close()
	}
	_, err = file.Write(content)
	if err != nil {
		fmt.Println("Error Writing file...", err)
		return false
	}
	return true
}
