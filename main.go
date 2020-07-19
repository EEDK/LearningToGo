package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	const loop = 10
	z := float64(1)

	for i := 0 ; i < loop ; i++{
		z = z - (z*z-x) / (2*z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
