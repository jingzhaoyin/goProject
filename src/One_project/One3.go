package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	n := len(digits)

	// 从最后一位开始遍历
	for i := n - 1; i >= 0; i-- {
		// 当前位加一
		digits[i]++

		// 如果加一后小于10，没有进位，直接返回
		if digits[i] < 10 {
			return digits
		}

		// 如果有进位，当前位设为0，继续处理前一位
		digits[i] = 0
	}

	// 如果所有位都处理完还有进位（如 999 -> 1000）
	// 需要在数组最前面添加1
	return append([]int{1}, digits...)
}

// 另一种实现方式，更直观的写法
func plusOne2(digits []int) []int {
	carry := 1 // 初始进位为1，相当于加一

	for i := len(digits) - 1; i >= 0 && carry > 0; i-- {
		sum := digits[i] + carry
		digits[i] = sum % 10
		carry = sum / 10
	}

	// 如果还有进位，在数组前面添加1
	if carry > 0 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

func main() {
	testCases := [][]int{
		{1, 2, 3},
		{4, 3, 2, 1},
		{9},
		{9, 9, 9},
		{0},
		{1, 9, 9},
		{8, 9, 9, 9},
	}

	fmt.Println("使用方法1：")
	for _, test := range testCases {
		fmt.Printf("%v -> %v\n", test, plusOne(test))
	}

	fmt.Println("\n使用方法2：")
	testCases2 := [][]int{
		{1, 2, 3},
		{4, 3, 2, 1},
		{9},
		{9, 9, 9},
		{0},
		{1, 9, 9},
		{8, 9, 9, 9},
	}
	for _, test := range testCases2 {
		fmt.Printf("%v -> %v\n", test, plusOne2(test))
	}
}
