package main

import (
	"fmt"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(10, 20)
}

func getCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	add := func(a, b float64) float64 {
		return a + b
	}

	f = func(c int) { return c }

	ff = func(fn func(int) int) {
		return fn(10)
	}

	fmt.Println("ff(f) = ", ff(f))

	fmt.Printf("%.2f\n", compute(add))

	counter := getCounter()
	counter2 := getCounter()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", counter())
		fmt.Printf("%d ", counter2())
	}
}
