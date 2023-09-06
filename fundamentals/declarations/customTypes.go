package declarations

import "fmt"

type Celcius float64
type Fahrenheit float64

func (c Celcius) CToF() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (f Fahrenheit) FToC() Celcius {
	return Celcius((f - 32) * 5 / 9)
}

func Cumpute() {
	var c Celcius = 18.25
	var f Fahrenheit = 84
	t := c + f.FToC()
	fmt.Println(t)
}
