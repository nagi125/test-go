package main

import "fmt"

// 無名関数
func main() {
	f := func(x int, y int) int {
		return x + y
	}

	i := f(1, 2)
	fmt.Println(i)
}