package main

import (
	"fmt"
	"sync"
	"time"
)

// 方法1：使用 WaitGroup 同步
func printNumbersWithWaitGroup() {
	fmt.Println("=== 方法1：使用 WaitGroup 同步 ===")

	var wg sync.WaitGroup

	// 启动奇数协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Print("奇数: ")
		for i := 1; i <= 10; i += 2 {
			fmt.Printf("奇数--%d ", i)
			time.Sleep(100 * time.Millisecond) // 模拟工作
		}
		fmt.Println()
	}()

	// 启动偶数协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Print("偶数: ")
		for i := 2; i <= 10; i += 2 {
			fmt.Printf("偶数--%d ", i)
			time.Sleep(10 * time.Millisecond) // 模拟工作
		}
		fmt.Println()
	}()

	// 等待所有协程完成
	wg.Wait()
	fmt.Println("两个协程都已完成")
}

// // 方法2：使用通道同步
// func printNumbersWithChannel() {
// 	fmt.Println("\n=== 方法2：使用通道同步 ===")

// 	// 创建通道用于同步
// 	done := make(chan bool, 2) // 缓冲通道，容量为2

// 	// 启动奇数协程
// 	go func() {
// 		fmt.Print("奇数: ")
// 		for i := 1; i <= 10; i += 2 {
// 			fmt.Printf("%d ", i)
// 			time.Sleep(50 * time.Millisecond)
// 		}
// 		fmt.Println()
// 		done <- true // 发送完成信号
// 	}()

// 	// 启动偶数协程
// 	go func() {
// 		fmt.Print("偶数: ")
// 		for i := 2; i <= 10; i += 2 {
// 			fmt.Printf("%d ", i)
// 			time.Sleep(50 * time.Millisecond)
// 		}
// 		fmt.Println()
// 		done <- true // 发送完成信号
// 	}()

// 	// 等待两个协程完成
// 	<-done
// 	<-done
// 	fmt.Println("两个协程都已完成")
// }

// // 方法3：交替打印奇偶数（更复杂的同步）
// func printNumbersAlternating() {
// 	fmt.Println("\n=== 方法3：交替打印奇偶数 ===")

// 	var wg sync.WaitGroup
// 	ch := make(chan bool, 1) // 用于控制打印顺序的通道

// 	wg.Add(2)

// 	// 奇数协程
// 	go func() {
// 		defer wg.Done()
// 		for i := 1; i <= 10; i += 2 {
// 			<-ch // 等待信号
// 			fmt.Printf("奇数: %d\n", i)
// 			ch <- true // 发送信号给偶数协程
// 		}
// 	}()

// 	// 偶数协程
// 	go func() {
// 		defer wg.Done()
// 		for i := 2; i <= 10; i += 2 {
// 			ch <- true // 发送信号给奇数协程
// 			<-ch       // 等待信号
// 			fmt.Printf("偶数: %d\n", i)
// 		}
// 	}()

// 	// 启动第一个协程
// 	ch <- true

// 	wg.Wait()
// 	close(ch)
// }

// // 方法4：使用互斥锁确保输出不交错
// func printNumbersWithMutex() {
// 	fmt.Println("\n=== 方法4：使用互斥锁确保完整输出 ===")

// 	var wg sync.WaitGroup
// 	var mu sync.Mutex

// 	wg.Add(2)

// 	// 奇数协程
// 	go func() {
// 		defer wg.Done()
// 		mu.Lock()
// 		defer mu.Unlock()

// 		fmt.Print("奇数: ")
// 		for i := 1; i <= 10; i += 2 {
// 			fmt.Printf("%d ", i)
// 			time.Sleep(100 * time.Millisecond)
// 		}
// 		fmt.Println()
// 	}()

// 	// 偶数协程
// 	go func() {
// 		defer wg.Done()

// 		// 等待一小段时间，让奇数协程先获取锁
// 		time.Sleep(10 * time.Millisecond)

// 		mu.Lock()
// 		defer mu.Unlock()

// 		fmt.Print("偶数: ")
// 		for i := 2; i <= 10; i += 2 {
// 			fmt.Printf("%d ", i)
// 			time.Sleep(100 * time.Millisecond)
// 		}
// 		fmt.Println()
// 	}()

// 	wg.Wait()
// }

// // 方法5：简单的并发打印（无同步）
// func printNumbersSimple() {
// 	fmt.Println("\n=== 方法5：简单并发（无同步）===")

// 	// 奇数协程
// 	go func() {
// 		fmt.Print("奇数: ")
// 		for i := 1; i <= 10; i += 2 {
// 			fmt.Printf("%d ", i)
// 		}
// 		fmt.Println()
// 	}()

// 	// 偶数协程
// 	go func() {
// 		fmt.Print("偶数: ")
// 		for i := 2; i <= 10; i += 2 {
// 			fmt.Printf("%d ", i)
// 		}
// 		fmt.Println()
// 	}()

// 	// 给协程一些时间执行
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("主程序结束")
// }

// // 方法6：使用 context 控制超时
// func printNumbersWithTimeout() {
// 	fmt.Println("\n=== 方法6：带超时的协程 ===")

// 	var wg sync.WaitGroup
// 	done := make(chan struct{})

// 	wg.Add(2)

// 	// 奇数协程
// 	go func() {
// 		defer wg.Done()
// 		fmt.Print("奇数: ")
// 		for i := 1; i <= 10; i += 2 {
// 			select {
// 			case <-done:
// 				return // 如果收到结束信号，立即返回
// 			default:
// 				fmt.Printf("%d ", i)
// 				time.Sleep(150 * time.Millisecond)
// 			}
// 		}
// 		fmt.Println()
// 	}()

// 	// 偶数协程
// 	go func() {
// 		defer wg.Done()
// 		fmt.Print("偶数: ")
// 		for i := 2; i <= 10; i += 2 {
// 			select {
// 			case <-done:
// 				return
// 			default:
// 				fmt.Printf("%d ", i)
// 				time.Sleep(150 * time.Millisecond)
// 			}
// 		}
// 		fmt.Println()
// 	}()

// 	// 设置超时
// 	go func() {
// 		time.Sleep(1 * time.Second) // 1秒后超时
// 		close(done)
// 	}()

// 	wg.Wait()
// 	fmt.Println("协程执行完成（可能超时）")
// }

func main() {
	// 运行各种方法的示例
	printNumbersWithWaitGroup()
	// printNumbersWithChannel()
	// printNumbersAlternating()
	// printNumbersWithMutex()
	// printNumbersSimple()
	// printNumbersWithTimeout()

	fmt.Println("\n=== 程序结束 ===")
}
