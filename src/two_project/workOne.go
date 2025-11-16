package main

import "fmt"

// 方式1：接收整数指针作为参数，直接修改指针指向的值
func addTen(ptr *int) {
	// 检查指针是否为 nil，避免空指针异常
	if ptr != nil {
		*ptr += 10 // 解引用并修改值
	}
}

// 方式2：另一种写法，更详细的注释
func addTenDetailed(ptr *int) {
	// 安全检查
	if ptr == nil {
		fmt.Println("错误：传入的指针为 nil")
		return
	}

	fmt.Printf("函数内 - 修改前: 指针地址 = %p, 指针指向的值 = %d\n", ptr, *ptr)

	// 通过指针修改值
	*ptr = *ptr + 10

	fmt.Printf("函数内 - 修改后: 指针地址 = %p, 指针指向的值 = %d\n", ptr, *ptr)
}

// 方式3：对比值传递（不会修改原值）
func addTenByValue(num int) {
	fmt.Printf("值传递函数内 - 修改前: 值 = %d\n", num)
	num += 10
	fmt.Printf("值传递函数内 - 修改后: 值 = %d\n", num)
	// 这里的修改不会影响外部的变量
}

func main() {
	fmt.Println("=== 指针传递演示 ===")

	// 示例1：基本指针使用
	num1 := 5
	fmt.Printf("主函数 - 调用前: num1 = %d, 地址 = %p\n", num1, &num1)

	addTen(&num1) // 传递 num1 的地址

	fmt.Printf("主函数 - 调用后: num1 = %d, 地址 = %p\n", num1, &num1)

	fmt.Println("\n=== 详细指针演示 ===")

	// 示例2：详细演示
	num2 := 15
	fmt.Printf("主函数 - 调用前: num2 = %d, 地址 = %p\n", num2, &num2)

	addTenDetailed(&num2)

	fmt.Printf("主函数 - 调用后: num2 = %d, 地址 = %p\n", num2, &num2)

	fmt.Println("\n=== 值传递对比演示 ===")

	// 示例3：对比值传递
	num3 := 25
	fmt.Printf("主函数 - 值传递调用前: num3 = %d, 地址 = %p\n", num3, &num3)

	addTenByValue(num3) // 传递值的副本

	fmt.Printf("主函数 - 值传递调用后: num3 = %d, 地址 = %p\n", num3, &num3)

	fmt.Println("\n=== 空指针安全演示 ===")

	// 示例4：空指针安全处理
	var nilPtr *int
	addTenDetailed(nilPtr) // 传入 nil 指针

	fmt.Println("\n=== 切片指针演示 ===")

	// 示例5：切片指针（切片本身是引用类型）
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("修改前: numbers = %v\n", numbers)

	modifySlice(numbers) // 切片不需要传递指针

	fmt.Printf("修改后: numbers = %v\n", numbers)
}

// 切片是引用类型，不需要指针
func modifySlice(slice []int) {
	for i := range slice {
		slice[i] += 10
	}
}
