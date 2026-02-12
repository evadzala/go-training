package main

import "fmt"

func main() {
	/*
		// 整數陣列
		var numbers [3]int
		numbers[0] = 15
		numbers[1] = 30
		numbers[2] = 18
		// numbers[3]=100 // 不能超過陣列的長度
		fmt.Println(numbers)
		fmt.Println(numbers[1] * 10)
		// 字串陣列
		var names [2]string = [2]string{"John", "Mary"}
		fmt.Println(names)
		// 取得陣列長度
		fmt.Println(len(numbers)) // 3
		fmt.Println(len(names))   // 2
	*/

	// 巡迴陣列中的資料
	var grades [4]int
	var index int
	// 逐一輸入陣列資料
	fmt.Println("請逐一輸入資料")
	for index = 0; index < len(grades); index++ {
		fmt.Scanln(&grades[index])
	}
	// 逐一取得陣列的資料
	var sum int = 0
	for index = 0; index < len(grades); index++ {
		sum += grades[index]
	}
	fmt.Println("輸入資料平均數為:", sum/len(grades))
}
