package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// 创建哈希表，存储数字和对应的索引
	numMap := make(map[int]int)

	for i, num := range nums {
		// 计算需要的补数
		complement := target - num

		// 检查补数是否在哈希表中
		if index, exists := numMap[complement]; exists {
			// 找到答案，返回两个索引
			return []int{index, i}
		}

		// 将当前数字和索引存入哈希表
		numMap[num] = i
	}

	// 如果没有找到，返回空切片
	return nil
}

// 暴力解法（不推荐，用于对比）
func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{3, 2, 4}, 6},
		{[]int{3, 3}, 6},
		{[]int{1, 2, 3, 4, 5}, 9},
		{[]int{1, 2, 3, 4, 5}, 10},
	}

	fmt.Println("使用哈希表解法：")
	for _, test := range testCases {
		result := twoSum(test.nums, test.target)
		fmt.Printf("nums: %v, target: %d -> %v\n", test.nums, test.target, result)
		if result != nil {
			fmt.Printf("  验证: nums[%d] + nums[%d] = %d + %d = %d\n",
				result[0], result[1], test.nums[result[0]], test.nums[result[1]], test.target)
		}
		fmt.Println()
	}

	// 性能对比
	fmt.Println("性能对比：")
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Printf("哈希表解法: %v\n", twoSum(nums, target))
	fmt.Printf("暴力解法: %v\n", twoSumBruteForce(nums, target))
}
