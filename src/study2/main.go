package main

import "fmt"

func pow(a [3]int) {

	for i, v := range a {
		a[i] = v * v
	}
	return
}

// スライスでは参照渡しとなる
func pow2(a []int) {
	for i, v := range a {
		a[i] = v * v
	}
}

// ポインタサンプル
func pow3(p *[3]int) {
	i := 0
	for i < 3 {
		(*p)[i] = (*p)[i] * (*p)[i]
		i++
	}
}

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

	// コピー
	s6 := []int{1, 2, 3, 4, 5}
	s7 := []int{10, 11}
	n  := copy(s6, s7)
	fmt.Println(n, s6, s7)

	// 値渡し aの値に変化はない
	a1 := [3]int{1, 2, 3}
	pow(a1)
	fmt.Println(a1)

	// 参照渡し aの値は変化する
	a2 := []int{1, 2, 3}
	pow2(a2)
	fmt.Println(a2)

	// 配列型とスライスの違い
	// スライスの場合は初期値がnilである
	var (
		a [3]int
		s []int
	)
	fmt.Println(a)
	fmt.Println(s == nil)

	// マップサンプル
	m1 := make(map[int]string)
	m1[1]  = "US"
	m1[81] = "Japan"
	m1[86] = "China"
	fmt.Println(m1)

	m2 := map[int]string{
		1: "Taro",
		2: "Hanako",
		3: "Jiro",
	}
	fmt.Println(m2)


	// ポインタサンプル
	var p *int
	fmt.Println(p == nil)

	p1 := &[3]int{1, 2, 3}
	pow3(p1)
	fmt.Println(p1)

}
