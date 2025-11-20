package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 使用 int64 类型的原子计数器
func withInt64Atomic() {
	fmt.Println("=== 使用 int64 原子计数器 ===")

	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
			fmt.Printf("协程 %d 完成\n", workerID)
		}(i)
	}

	wg.Wait()
	fmt.Printf("最终计数器值: %d\n", counter)

	if counter == 10000 {
		fmt.Println("✓ int64 原子操作测试通过")
	}
}

func main() {
	// 运行基础版本
	fmt.Println("=== atomic 原子操作实现无锁计数器 ===\n")

	var counter int32
	var wg sync.WaitGroup

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&counter, 1)
			}
		}(i)
	}

	wg.Wait()

	// 输出结果
	fmt.Printf("最终计数器值: %d\n", counter)
	fmt.Printf("期望值: %d\n", 10*1000)

	if counter == 10000 {
		fmt.Println("✓ 原子操作成功实现无锁计数器")
	}

	// 运行64位版本
	withInt64Atomic()
}
