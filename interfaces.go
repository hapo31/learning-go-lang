package main

import "fmt"

type A struct {
	x, y int
}

type B struct {
	a, b int
}

type Addable interface {
	add(int, int)
}

func (target *A) add(x, y int) {
	target.x += x
	target.y += y
}

func (target *B) add(a, b int) {
	target.a += a
	target.b += b
}

func Add(target Addable, val1, val2 int) {
	target.add(val1, val2)
}

func main() {
	a := A{10, 10}
	b := B{20, 20}
	fmt.Printf("%v %v\n", a, b)
	Add(&a, 5, 2) // 型がinterfaceの場合はポインタを渡す必要がある
	Add(&b, 10, 4)
	fmt.Printf("%v %v\n", a, b)

}
