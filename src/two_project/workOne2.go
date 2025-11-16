package main

import "fmt"

// 方式1：接收切片指针，修改切片中的每个元素
func doubleSliceElements(slicePtr *[]int) {
	// 安全检查
	if slicePtr == nil {
		fmt.Println("错误：切片指针为 nil")
		return
	}

	if *slicePtr == nil {
		fmt.Println("警告：切片为 nil")
		return
	}

	// 通过指针访问切片并修改元素
	slice := *slicePtr
	for i := 0; i < len(slice); i++ {
		slice[i] = slice[i] * 2
	}
}

// 方式2：使用 range 循环（更简洁）
func doubleSliceElementsWithRange(slicePtr *[]int) {
	if slicePtr == nil || *slicePtr == nil {
		return
	}

	slice := *slicePtr
	for i := range slice {
		slice[i] *= 2 // 简写形式
	}
}

// 方式3：对比 - 直接传递切片（不需要指针，因为切片本身就是引用类型）
func doubleSliceDirectly(slice []int) {
	for i := range slice {
		slice[i] *= 2
	}
}

// 方式4：创建新切片并返回（不修改原切片）
func doubleSliceAndReturn(slice []int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = v * 2
	}
	return result
}

// 方式5：修改切片本身（如追加元素）
func doubleAndAppend(slicePtr *[]int) {
	if slicePtr == nil {
		return
	}

	// 如果需要修改切片本身（长度、容量），需要指针
	original := *slicePtr
	for i := range original {
		original[i] *= 2
	}

	// 追加新元素到切片
	*slicePtr = append(original, 100, 200)
}

func main() {
	fmt.Println("=== 方式1：使用切片指针修改元素 ===")
	numbers1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("修改前: %v, 长度: %d, 容量: %d\n", numbers1, len(numbers1), cap(numbers1))

	doubleSliceElements(&numbers1)
	fmt.Printf("修改后: %v, 长度: %d, 容量: %d\n", numbers1, len(numbers1), cap(numbers1))

	fmt.Println("\n=== 方式2：使用 range 循环 ===")
	numbers2 := []int{10, 20, 30}
	fmt.Printf("修改前: %v\n", numbers2)

	doubleSliceElementsWithRange(&numbers2)
	fmt.Printf("修改后: %v\n", numbers2)

	fmt.Println("\n=== 方式3：直接传递切片（对比）===")
	numbers3 := []int{5, 10, 15}
	fmt.Printf("修改前: %v\n", numbers3)

	doubleSliceDirectly(numbers3) // 不需要取地址
	fmt.Printf("修改后: %v\n", numbers3)

	fmt.Println("\n=== 方式4：返回新切片（不修改原切片）===")
	numbers4 := []int{1, 3, 5}
	fmt.Printf("原切片: %v\n", numbers4)

	newSlice := doubleSliceAndReturn(numbers4)
	fmt.Printf("新切片: %v\n", newSlice)
	fmt.Printf("原切片未改变: %v\n", numbers4)

	fmt.Println("\n=== 方式5：修改切片本身（需要指针）===")
	numbers5 := []int{2, 4, 6}
	fmt.Printf("修改前: %v, 长度: %d, 容量: %d\n", numbers5, len(numbers5), cap(numbers5))

	doubleAndAppend(&numbers5)
	fmt.Printf("修改后: %v, 长度: %d, 容量: %d\n", numbers5, len(numbers5), cap(numbers5))

	fmt.Println("\n=== 边界情况测试 ===")

	// 测试空切片
	emptySlice := []int{}
	fmt.Printf("空切片修改前: %v\n", emptySlice)
	doubleSliceElements(&emptySlice)
	fmt.Printf("空切片修改后: %v\n", emptySlice)

	// 测试 nil 切片
	var nilSlice []int
	fmt.Printf("nil切片修改前: %v, 是否为nil: %t\n", nilSlice, nilSlice == nil)
	doubleSliceElements(&nilSlice)
	fmt.Printf("nil切片修改后: %v, 是否为nil: %t\n", nilSlice, nilSlice == nil)

	// 测试 nil 指针
	var nilPtr *[]int
	doubleSliceElements(nilPtr) // 应该输出错误信息
}

// 额外示例：处理多维切片
func double2DSlice(slicePtr *[][]int) {
	if slicePtr == nil {
		return
	}

	slice2D := *slicePtr
	for i := range slice2D {
		for j := range slice2D[i] {
			slice2D[i][j] *= 2
		}
	}
}

func demonstrate2DSlice() {
	fmt.Println("\n=== 多维切片示例 ===")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Printf("修改前: %v\n", matrix)
	double2DSlice(&matrix)
	fmt.Printf("修改后: %v\n", matrix)
}
