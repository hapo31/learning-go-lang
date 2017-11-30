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
	// レシーバーは実装でnilガードをしておくとよさそう
	if target != nil {
		target.x += x
		target.y += y
	}
}

func (target *B) add(a, b int) {
	if target != nil {
		target.a += a
		target.b += b
	} else {
		fmt.Println("<nil>")
	}
}

func Add(target Addable, val1, val2 int) {
	target.add(val1, val2)
}

func main() {
	a := A{10, 10}
	b := B{20, 20}

	var nilAdder Addable
	var nilB *B
	nilAdder = nilB

	fmt.Printf("%v %v\n", a, b)
	Add(&a, 5, 2) // レシーバーで要求している型がポインタなのでポインタを渡す必要がある
	Add(&b, 10, 4)
	Add(nilAdder, 5, 5)

	fmt.Printf("%v %v\n", a, b)

	PrintTypeInfo(a)
	PrintTypeInfo(b)

	// createPointer内で作ったオブジェクトのポインタを受け取って、変数の寿命はどうなるか見てみる
	var x = createPointer()
	// どうやらスコープを抜けることで寿命が切れてポインタ先が消えることはないらしい(参照が維持されている？)
	PrintTypeInfo(x)
	Add(x, 10, 10)
	PrintTypeInfo(x)
}

type T interface {
}

func PrintTypeInfo(t T) {
	fmt.Printf("%T [%v]\n", t, t)
}

// スコープ内でオブジェクト作ってポインタを返してみる
func createPointer() *A {
	return &A{10, 10}
}
