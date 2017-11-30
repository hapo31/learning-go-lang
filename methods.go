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

// ポインタレシーバーの定義
// ポインタを受け取るので、元のオブジェクトを書き換えられる
func (p *Pos) Move(x, y int) {
	p.x = x
	p.y = y
}

func main() {
	p := Pos{10, 10}
	fmt.Printf("%d, %d\n", p.x, p.y) // レシーバーの使い方

	p.Move(15, 15)
	fmt.Printf("%d, %d\n", p.x, p.y) // ポインタレシーバーの使い方

}
