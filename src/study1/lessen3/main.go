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



	// 浮動小数点型
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	// 自動で64bitになる
	fl := 3.2
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n", fl64, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)

	// 型変換
	fmt.Printf("%T\n", float64(fl32))

	// 無限
	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	// -無限
	ninf := -1.0 / zero
	fmt.Println(ninf)

	// 非数
	nan := zero / zero
	fmt.Println(nan)



	// uint型
	var u8 uint = 255
	fmt.Println(u8)



	// 複素数型
	var c64 complex64 = -5 + 12i
	fmt.Println(c64)



	// 論理値型
	var t, f bool = true, false
	fmt.Println(t, f)

	


}