package main

import (
	"fmt"
	"strconv"
)

func main() {
	// エラーハンドリング
	var s string = "100"
	//var s string = "A"

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("i = %T\n", i)


	// 繰り返し処理
	/*
	point := 0
	for point < 10 {
		fmt.Println("Loop")
		point++
	}
	*/

	// よくあるfor文
	/*
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}

		if i == 6 {
			break
		}
		fmt.Println(i)
	}
	*/

	// 配列を使ったfor
	/*
	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	*/

	// Rangeを使ったfor
	/*
	arr := [3]int{1, 2, 3}
	for i, v := range arr {
		fmt.Println(i, v)
	}

	sl := []string{"Python", "PHP", "Go"}
	for i, v := range sl {
		fmt.Println(i, v)
	}
	*/

	// 連想配列 for
	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}

}
