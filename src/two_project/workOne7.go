package main

import (
	"fmt"
	"time"
)

// 最简洁的实现
func simpleProducerConsumer() {
	ch := make(chan int)

	// 生产者
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	// 消费者
	for num := range ch {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

func main() {
	fmt.Println("最简单的生产-消费示例:")
	simpleProducerConsumer()

	// 或者使用更清晰的方式
	fmt.Println("\n清晰版本:")
	ch := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Printf("发送: %d\n", i)
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(ch)
	}()

	for num := range ch {
		fmt.Printf("接收: %d\n", num)
	}
}
