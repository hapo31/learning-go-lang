package main

import (
	"fmt"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(10, 20)
}

func main() {
	add := func(a, b float64) {
		return a + b
	}

	fmt.Printf("%.2f", compute(add))
}
