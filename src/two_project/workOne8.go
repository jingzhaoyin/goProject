package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== 带缓冲通道的生产者-消费者示例 ===")

	// 创建缓冲大小为10的通道
	ch := make(chan int, 10)

	// 使用WaitGroup等待所有协程完成
	var wg sync.WaitGroup

	// 启动生产者协程
	wg.Add(1)
	go producer(ch, &wg)

	// 启动消费者协程
	wg.Add(1)
	go consumer(ch, &wg)

	// 等待所有协程完成
	wg.Wait()
	fmt.Println("程序执行完成")
}

// 生产者：向通道发送100个整数
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch) // 关闭通道，通知消费者没有更多数据

	fmt.Println("生产者开始工作...")

	for i := 1; i <= 100; i++ {
		ch <- i // 发送数据到通道
		fmt.Printf("生产者发送: %d (通道长度: %d/%d)\n", i, len(ch), cap(ch))
		time.Sleep(10 * time.Millisecond) // 模拟生产耗时
	}

	fmt.Println("生产者完成工作，已关闭通道")
}

// 消费者：从通道接收整数并打印
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("消费者开始工作...")
	count := 0

	for num := range ch { // 循环接收，直到通道关闭
		count++
		fmt.Printf("消费者接收: %d (总计: %d, 通道长度: %d/%d)\n",
			num, count, len(ch), cap(ch))
		time.Sleep(15 * time.Millisecond) // 模拟消费耗时
	}

	fmt.Printf("消费者完成工作，共处理 %d 个数据\n", count)
}
