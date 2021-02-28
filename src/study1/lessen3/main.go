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



	// 文字列型
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	// 複数行表現
	fmt.Println(`test
	test
		test`)

	// "を文字列として扱いたい場合
	fmt.Println("\"")
	fmt.Println(`"`)

	// 文字列はByte配列の集まり
	fmt.Println(string(s[0]))



	// byte型
	byteA := []byte{72, 73}
	fmt.Println(byteA);

	// 文字列に変換
	fmt.Println(string(byteA))

	// 文字列をbyteの配列に変換
	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))



	// 配列型(個数もセットで型になっている)
	// 可変長にしたい場合は後述のスライスを利用する
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	// 要素数を自動でカウントしてくれる
	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	// 値の取り出し
	fmt.Println(arr2[0])

	// 配列の更新
	arr2[2] = "C"
	fmt.Println(arr2)

	// 要素数が違う場合は違う型となりエラーとなる
	// var arr5 [4]int
	// arr5 = arr1

	// 要素数を調べる
	fmt.Println(len(arr1))



}