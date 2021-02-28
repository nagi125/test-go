package main

import "fmt"

func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

// クロージャーの実装
func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

// ジェネレーターの実装
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
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


	// クロージャー
	f3 := Later()
	fmt.Println(f3("Hello"))


	// ジェネレーター
	ints := integers()

	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	otherInts := integers()
	fmt.Println(otherInts())
	fmt.Println(otherInts())
	fmt.Println(otherInts())

}