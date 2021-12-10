package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	Freezing      Celsius = 0
	BoilingC      Celsius = 100

	FreezingK Kelvin = 273
	BoilingK  Kelvin = 373.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g ºC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g ºF", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g ºK", k) }
