package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	// 负数不是回文数
	if x < 0 {
		return false
	}

	// 将整数转换为字符串
	s := strconv.Itoa(x)

	// 使用双指针法判断是否为回文
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// 另一种方法：不转换为字符串，通过数学运算
func isPalindrome2(x int) bool {
	// 负数和以0结尾的非0数不是回文数
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	// 特殊情况：0是回文数
	if x == 0 {
		return true
	}

	// 反转后半部分数字
	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 当数字长度为奇数时，通过 revertedNumber/10 去除中间位
	return x == revertedNumber || x == revertedNumber/10
}

func main() {
	testCases := []int{121, -121, 10, 0, 12321, 12345}

	fmt.Println("使用方法1（字符串转换）：")
	for _, num := range testCases {
		fmt.Printf("%d: %t\n", num, isPalindrome(num))
	}

	fmt.Println("\n使用方法2（数学运算）：")
	for _, num := range testCases {
		fmt.Printf("%d: %t\n", num, isPalindrome2(num))
	}
}
