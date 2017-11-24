package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello world.", f(3))
	v := rand.Intn(10)

	// pointer
	pv := &v

	// switch
	switch {
	case v < 10:
		fmt.Println("less than 10")
	case v >= 10:
		fmt.Println("more than 10")
	default:
		fmt.Println("many")
	}

	fmt.Println("v Pointer", pv)
	fmt.Println("v derefference", *pv)

	Fibo(100)
	defer (func() {
		fmt.Println("end of process.")
	})()
}

func f(n int) (r int) {
	r = 1
	if n != 0 {
		return f(n-1) * n
	}
	return
}

func Fibo(max int) {
	var now = 1
	var prev = 0
	fmt.Print(now, " ")
	for now < max {
		result := now + prev
		fmt.Print(result, " ")
		prev = now
		now = result
	}
}

// $ go run main.go
// $ go build main.go & main.exe
