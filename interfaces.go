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

	var interfaceVal interface{} = 10
	// interfaceの型アサーション
	var c1, ok1 = interfaceVal.(string)
	fmt.Println(c1, ok1)

	var c2, ok2 = interfaceVal.(float64)
	fmt.Println(c2, ok2)

	// switch文で一気にアサーションと分岐
	switch v := interfaceVal.(type) {
	case int:
		fmt.Printf("int: v * 2 = %d\n", v*2)
	case string:
		fmt.Println("string:", v)
	case float64:
		fmt.Printf("float: %.2f", v)
	}

	var person = Person{"マイケル☆", 114514}
	fmt.Println(person)
}

type Person struct {
	Name string
	Age  int
}

// 独自型の文字列化はString()というレシーバーを定義して行う
func (p Person) String() string {
	return fmt.Sprintf("Name:%v  Age:%v", p.Name, p.Age)
}

// 空のインターフェースは任意の値を渡せる
type T interface {
}

func PrintTypeInfo(t T) { // ここはTの代わりにinterface{}でもいいらしい
	fmt.Printf("%T [%v]\n", t, t)
}

// スコープ内でオブジェクト作ってポインタを返してみる
func createPointer() *A {
	return &A{10, 10}
}
