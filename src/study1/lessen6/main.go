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
}
