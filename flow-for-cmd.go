package main

import "fmt"

func main() {
	// break
	/*
		var x int = 0
		for x < 5 {
			if x == 3 {
				break
			}
			fmt.Println(x)
			x++
		}
	*/

	// continue
	/*
		var x int
		for x = 0; x < 5; x++ {
			if x%2 == 0 {
				continue
			}
			fmt.Println(x)
		}
	*/
	// 實際範例，持續讓使用者輸入數字，計算總合，直到使用者輸入0為止
	var result int = 0
	for true { // 無窮迴圈
		var n int
		fmt.Scanln(&n)
		if n == 0 {
			break
		}
		result += n
	}
	fmt.Println("總和", result)
}
