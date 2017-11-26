package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello world.", f(3))
	// 宣言して代入
	v := rand.Intn(10)
	// 使われてないとエラー
	// var f = 10

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

	rect := Rect{10, 10, 20, 20}
	rectp := &rect

	// 構造体をコピーすると参照ではなく値のコピーになる
	rectCopy := rect
	fmt.Println("before", rectCopy)
	// 元のインスタンスの値を書き換えてもコピー先には影響なし
	rect.top = 100
	fmt.Println("after", rectCopy)

	fmt.Println("rect:", rect.top)
	// ポインタの参照先もドットアクセス可
	fmt.Println("rect p ", rectp.right)
	// Cっぽくも書ける
	fmt.Println((*rectp).top)

	// 配列
	// 静的？に確保される
	var arr [5]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	// invalid array index 5 (out of bounds for 5-element array)
	// arr[5] = 6

	fmt.Println("array:", arr)
	// スライス
	var slice = arr[1:4]
	fmt.Println("slice", slice)
	// スライスは配列への参照を保持するので、もとの値を変更するとsliceから参照する値も変更される
	arr[3] = 10
	fmt.Println("slice", slice)

	// 配列と初期化
	var arr1 = [5]bool{true, true, false, false, false}
	// スライスの初期化
	var slice1 = []bool{true, true, false, false, false}

	fmt.Println(arr1, slice1)
	fmt.Println(slice1[0:2]) // スライスの作成 [begin, end)
	slice1 = slice1[:2]
	printSliceInfo(slice1)
	slice1 = slice1[0:5] // 縮めたやつを戻すと復活するが、省略記法を使うと縮めたあとの長さのままになる
	printSliceInfo(slice1)

	// 遅延実行。関数を抜ける前に実行される
	defer (func() {
		fmt.Println("\n---------------------------\nend of process.")
	})()
}

// 関数
func f(n int) (r int) { // 戻り値を表す変数にここで名前を付けられる
	r = 1
	if n != 0 {
		return f(n-1) * n
	}
	return // 戻り値に名前をつけた場合はreturnだけで値が返されるようになる
}

// 関数
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

func printSliceInfo(s []bool) {
	// sliceの情報はlenとcapでそれぞれ取れる
	fmt.Printf("%v size:%d cap:%d\n", s, len(s), cap(s))
}

// 構造体
type Rect struct {
	left   int
	right  int
	top    int
	bottom int
}

// $ go run main.go
// $ go build main.go & main.exe
