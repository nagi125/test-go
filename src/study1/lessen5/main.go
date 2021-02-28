package main

import "fmt"

func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func main() {
	// 無名関数
	f := func(x int, y int) int {
		return x + y
	}

	i := f(1, 2)
	fmt.Println(i)


	// 関数を返す関数
	f2 := ReturnFunc()
	f2()

}