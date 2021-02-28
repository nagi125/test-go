package main

import "fmt"

// 型
func main() {
	var i int = 100

	var i2 int64 = 200

	fmt.Println(i + 50)

	// 型が違う場合は計算できない
	// fmt.Println(i + i2)

	// 型を調べる
	fmt.Printf("%T\n", i2)

	// 型変換
	fmt.Printf("%T\n", int32(i2))

	// 型を揃えれば計算可能
	fmt.Println(i + int(i2))

}