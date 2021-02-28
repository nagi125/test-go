package main

import "fmt"

// 定数
// 頭文字を大文字にすると別のパッケージからも呼び出せる
const Pi = 3.14

const (
	URL      = "http://xxx.co.jp"
	SiteName = "test"
)

// 同じ操作の場合は省略可能
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

func main() {
	fmt.Println(Pi)

	// 定数は上書きできない
	// Pi = 3
	// fmt.Println(Pi)

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)

}