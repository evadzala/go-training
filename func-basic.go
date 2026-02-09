package main

import "fmt"

func sum(max int) {
	var result int = 0
	var n int
	for n = 1; n <= max; n++ {
		result += n
	}
	fmt.Println(result)
}
func main() {
	// 計算1+2+3+...+max的函式包裝
	sum(100)
	sum(50)
	sum(10)
}
