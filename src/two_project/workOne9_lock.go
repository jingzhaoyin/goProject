package main

import (
	"fmt"
	"sync"
)

func main() {
	// 共享计数器
	var counter int

	// 互斥锁，用于保护共享计数器
	var mutex sync.Mutex

	// 等待组，用于等待所有协程完成
	var wg sync.WaitGroup

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1) // 每启动一个协程，等待组加1

		go func(workerID int) {
			defer wg.Done() // 协程结束时通知等待组

			// 每个协程执行1000次递增操作
			for j := 0; j < 1000; j++ {
				// 在修改共享数据前获取锁
				mutex.Lock()

				// 临界区：对共享计数器进行递增操作
				counter++

				// 操作完成后释放锁
				mutex.Unlock()
			}

			fmt.Printf("协程 %d 完成1000次递增操作\n", workerID)
		}(i) // 将循环变量作为参数传入，避免闭包问题
	}

	// 等待所有协程执行完毕
	wg.Wait()

	// 输出最终结果
	fmt.Printf("\n最终计数器值: %d\n", counter)
	fmt.Printf("期望值: %d\n", 10*1000)

	// 验证结果是否正确
	if counter == 10000 {
		fmt.Println("✓ 结果正确：计数器得到正确保护，无竞态条件")
	} else {
		fmt.Println("✗ 结果错误：存在竞态条件，数据不一致")
	}
}
