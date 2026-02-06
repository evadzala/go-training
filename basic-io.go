package main

import "fmt"

func main() {
	// 基本輸入練習: fmt.Println(資料,資料,...)
	// fmt.Println(3, "Hello", true)

	// 基本輸入練習: fmt.Scanln(&變數名稱, &變數名稱, ...)
	// &變數名稱: 取得變數的指標(Pointer)
	/*
		var x int
		fmt.Print("輸入一個數字")
		fmt.Scanln(&x)
		fmt.Println(x)
	*/

	// 整合練習: 輸入兩個字，並輸出數字相加的結果
	/*
		var x int
		var y int
		fmt.Print("輸入第一個數字:")
		fmt.Scanln(&x)
		fmt.Print("輸入第二個數字:")
		fmt.Scanln(&y)
		var result int = x + y
		fmt.Println(result)
	*/

	var x int
	var y int
	fmt.Print("輸入兩個數字，用空格隔開:")
	fmt.Scanln(&x, &y)
	var result int = x + y
	fmt.Println(result)
}
