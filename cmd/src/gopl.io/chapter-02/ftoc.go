// ftoc shows the boiling and freezing points of water in degrees F and C
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // 32°F = 0°C
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))
}

// converts a temperature in F to C
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
