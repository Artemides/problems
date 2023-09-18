package interfaces

import (
	"flag"
	"fmt"
)

type Celcius float64
type Fahrenheit float64

func (c Celcius) CToF() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (f Fahrenheit) FToC() Celcius {
	return Celcius((f - 32) * 5 / 9)
}

type celsiusFlag struct {
	Celcius
}

func (c *celsiusFlag) Set(str string) error {
	var unit string
	var value float64
	fmt.Sscanf(str, "%f%s", &value, &unit)
	if unit == "C" || unit == "°C" {
		c.Celcius = Celcius(value)
		return nil
	}

	if unit == "F" || unit == "°F" {
		c.Celcius = (Fahrenheit(value)).FToC()
		return nil
	}

	return fmt.Errorf("invalid temperature %v", c)
}

func (c celsiusFlag) String() string {
	return fmt.Sprintf("%3.f°C", c.Celcius)
}

func CelsiusFlag(name string, value Celcius, usage string) *Celcius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celcius
}

func RunTemperature() {
	var temp = CelsiusFlag("temp", 20.0, "the temperature")
	flag.Parse()
	fmt.Printf("%2.fC", *temp)
}
