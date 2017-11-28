// https://go-tour-jp.appspot.com/moretypes/26

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var prev = 0
	var now = 0
	return func() (r int) {
		r = prev + now
		prev = now
		now = r
		if r == 0 {
			prev = 1
		}
		return
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
