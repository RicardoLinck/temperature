package tempconv

import "fmt"

// Celsius represent the temperature in Celsius
type Celsius float64

// Fahrenheit represent the temperature in Fahrenheit
type Fahrenheit float64

// Kelvin represent the temperature in Kelvin
type Kelvin float64

const (
	// AbsoluteZeroC Absolute Zero temperature in celsius
	AbsoluteZeroC Celsius = -273.15
	// FreezingC Freezing temperatuer in celsius
	FreezingC Celsius = 0
	// BoilingC Boliing temperature in celsius
	BoilingC Celsius = 100
)

// CtoF converts celsius to fahrenheit
func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FtoC converts fahrenheit to celsius
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KtoC converts kelvin to celsius
func KtoC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// CtoK converts celsius to kelvin
func CtoK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
