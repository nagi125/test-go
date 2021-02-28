package main

import "fmt"

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	// 変数の型だけ宣言すると、型の初期値が入る
	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	// 値が更新される
	i = 150
	fmt.Println(i)

	// 暗黙的な定義
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// 再定義できない
	// i4 := 500

	// 最初と型が違う場合は代入できない
	// i4 = "Hello"
	//fmt.Println(i4)

	// 関数呼び出し
	outer()
}