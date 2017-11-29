package main

import (
	"fmt"
)

type Pos struct {
	x, y int
}

// レシーバーの定義
func (p Pos) Index(width int) int {
	return p.y*width + p.x
}

func main() {
	p := Pos{10, 10}
	fmt.Printf("%d\n", p.Index(5)) // レシーバーの使い方
}
