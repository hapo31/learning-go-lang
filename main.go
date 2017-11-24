package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world.", f(3))
	Fibo(20)
}

func f(n int) (r int) {
	r = 1
	if n != 0 {
		return f(n-1) * n
	}
	return
}

func Fibo(max int) {
	var result = 1
	var prev = 0

	for ; result < max; prev = result {
		result += prev
		fmt.Println(result)
	}
}

// $ go run main.go
// $ go build main.go & main.exe
