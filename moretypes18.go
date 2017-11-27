// https://go-tour-jp.appspot.com/moretypes/18 で書いたやつ↓

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) (result [][]uint8) {
	result = make([][]uint8, dy, dy)
	for y := range result {
		result[y] = make([]uint8, dx, dx)
		for x := range result[y] {
			result[y][x] = uint8(x*x + y*y)
		}
	}
	return
}

func main() {
	pic.Show(Pic)
}
