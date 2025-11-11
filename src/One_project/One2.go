package main

import (
	"fmt"
)

func isValid(s string) bool {
	// 使用切片模拟栈
	stack := make([]rune, 0)

	// 定义括号映射关系
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// 遍历字符串中的每个字符
	for _, char := range s {
		// 如果是左括号，压入栈中
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			// 如果是右括号，检查栈是否为空
			if len(stack) == 0 {
				return false
			}

			// 弹出栈顶元素并检查是否匹配
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if top != bracketMap[char] {
				return false
			}
		}
	}

	// 最后栈应该为空
	return len(stack) == 0
}

// 另一种实现方式，使用 switch 语句
func isValid2(s string) bool {
	stack := make([]rune, 0)

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			// 左括号入栈
			stack = append(stack, char)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	testCases := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
		"",
		"(",
		")",
	}

	fmt.Println("使用方法1（映射表）：")
	for _, test := range testCases {
		fmt.Printf("\"%s\": %t\n", test, isValid(test))
	}

	fmt.Println("\n使用方法2（switch语句）：")
	for _, test := range testCases {
		fmt.Printf("\"%s\": %t\n", test, isValid2(test))
	}
}
