package main

import "fmt"

func main() {

	// スライスサンプル
	// 要素10
	s1 := make([]int, 10)
	fmt.Println(s1)
	fmt.Println(len(s1))

	// 要素5 容量10
	s2 := make([]int, 5, 10)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	// 簡易スライス
	s3 := [5]int{1, 2, 3, 4, 5}
	s4 := s3[0:2]
	fmt.Println(s4)

	// 要素の追加
	s5 := []int{1, 2, 3}
	s5 = append(s5, 4)
	fmt.Println(s5)
}
