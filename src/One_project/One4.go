package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 慢指针，指向最后一个不重复元素的位置
	i := 0

	// 快指针，遍历整个数组
	for j := 1; j < len(nums); j++ {
		// 当遇到不相同的元素时
		if nums[j] != nums[i] {
			// 移动慢指针
			i++
			// 将不重复元素复制到正确位置
			nums[i] = nums[j]
		}
	}

	// 返回新长度（索引+1）
	return i + 1
}

// 另一种写法，更详细的注释版本
func removeDuplicates2(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	// i 指向当前已处理的不重复部分的最后一个位置
	i := 0

	// j 从第二个元素开始遍历
	for j := 1; j < n; j++ {
		// 如果当前元素与前一个不重复元素不同
		if nums[j] != nums[i] {
			// 先移动指针，再赋值
			i++
			nums[i] = nums[j]
		}
		// 如果相同，j 继续向后移动，i 保持不变
	}

	return i + 1
}

func main() {
	testCases := [][]int{
		{1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		{1, 2, 3},
		{1, 1, 1, 1},
		{},
		{1},
	}

	for _, nums := range testCases {
		// 复制原数组，避免修改影响后续测试
		testNums := make([]int, len(nums))
		copy(testNums, nums)

		newLength := removeDuplicates(testNums)

		fmt.Printf("原数组: %v\n", nums)
		fmt.Printf("新长度: %d\n", newLength)
		fmt.Printf("处理后的数组: %v\n", testNums[:newLength])
		fmt.Println("---")
	}
}
