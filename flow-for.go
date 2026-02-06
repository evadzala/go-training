package main

import "fmt"

func main() {
	// 基本迴圈使用
	/*
		var x int = 0
		for x < 3 {
			fmt.Println("Hello")
			x++
		}
	*/
	/*
		var x int
		for x = 0; x < 5; x += 2 {
			fmt.Println(x)
		}
	*/

	// 實際範例, 計算1+2+3+...+50的結果
	var result int = 0
	var x int = 1
	var y int
	fmt.Println("要從1累加到多少?")
	fmt.Scanln(&y)
	for x <= y {
		result = result + x
		x++
	}
	fmt.Println(result)
}
