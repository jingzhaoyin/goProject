package main

import (
	"fmt"
	"sync"
	"time"
)

// 任务类型定义
type Task func() error

// 任务结果
type TaskResult struct {
	TaskID    int
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
	Error     error
	Success   bool
}

// 任务调度器
type TaskScheduler struct {
	tasks      []Task
	results    []TaskResult
	mu         sync.Mutex
	wg         sync.WaitGroup
	maxWorkers int
	timeout    time.Duration
}

// 新建任务调度器
func NewTaskScheduler(maxWorkers int) *TaskScheduler {
	return &TaskScheduler{
		tasks:      make([]Task, 0),
		results:    make([]TaskResult, 0),
		maxWorkers: maxWorkers,
	}
}

// 设置超时时间
func (ts *TaskScheduler) SetTimeout(timeout time.Duration) {
	ts.timeout = timeout
}

// 添加任务
func (ts *TaskScheduler) AddTask(task Task) {
	ts.tasks = append(ts.tasks, task)
}

// 添加多个任务
func (ts *TaskScheduler) AddTasks(tasks []Task) {
	ts.tasks = append(ts.tasks, tasks...)
}

// 串行执行任务
func (ts *TaskScheduler) RunSerial() {
	fmt.Println("=== 串行执行任务 ===")
	ts.results = make([]TaskResult, len(ts.tasks))

	for i, task := range ts.tasks {
		result := ts.executeTask(i, task)
		ts.results[i] = result
		ts.printTaskResult(result)
	}

	ts.printSummary()
}

// 并行执行任务（使用工作池）
func (ts *TaskScheduler) RunParallel() {
	fmt.Println("=== 并行执行任务（工作池模式）===")
	taskCount := len(ts.tasks)
	ts.results = make([]TaskResult, taskCount)

	// 创建任务通道
	taskChan := make(chan int, taskCount)
	resultChan := make(chan TaskResult, taskCount)

	// 启动工作协程
	for i := 0; i < ts.maxWorkers && i < taskCount; i++ {
		ts.wg.Add(1)
		go ts.worker(taskChan, resultChan)
	}

	// 发送任务到通道
	for i := range ts.tasks {
		taskChan <- i
	}
	close(taskChan)

	// 收集结果
	go func() {
		ts.wg.Wait()
		close(resultChan)
	}()

	// 处理结果
	for result := range resultChan {
		ts.results[result.TaskID] = result
		ts.printTaskResult(result)
	}

	ts.printSummary()
}

// 工作协程
func (ts *TaskScheduler) worker(taskChan <-chan int, resultChan chan<- TaskResult) {
	defer ts.wg.Done()

	for taskID := range taskChan {
		result := ts.executeTask(taskID, ts.tasks[taskID])
		resultChan <- result
	}
}

// 执行单个任务
func (ts *TaskScheduler) executeTask(taskID int, task Task) TaskResult {
	startTime := time.Now()
	var err error
	var success bool

	if ts.timeout > 0 {
		// 带超时执行
		err = ts.executeWithTimeout(task)
	} else {
		// 普通执行
		err = task()
	}

	endTime := time.Now()

	if err == nil {
		success = true
	}

	return TaskResult{
		TaskID:    taskID,
		StartTime: startTime,
		EndTime:   endTime,
		Duration:  endTime.Sub(startTime),
		Error:     err,
		Success:   success,
	}
}

// 带超时执行任务
func (ts *TaskScheduler) executeWithTimeout(task Task) error {
	done := make(chan error, 1)

	go func() {
		done <- task()
	}()

	select {
	case err := <-done:
		return err
	case <-time.After(ts.timeout):
		return fmt.Errorf("任务执行超时")
	}
}

// 打印单个任务结果
func (ts *TaskScheduler) printTaskResult(result TaskResult) {
	status := "✓ 成功"
	if !result.Success {
		status = "✗ 失败"
	}

	fmt.Printf("任务 %d [%s] - 耗时: %v",
		result.TaskID+1, status, result.Duration)

	if result.Error != nil {
		fmt.Printf(" - 错误: %v", result.Error)
	}
	fmt.Println()
}

// 打印执行摘要
func (ts *TaskScheduler) printSummary() {
	var totalDuration time.Duration
	successCount := 0
	failureCount := 0

	for _, result := range ts.results {
		totalDuration += result.Duration
		if result.Success {
			successCount++
		} else {
			failureCount++
		}
	}

	avgDuration := totalDuration / time.Duration(len(ts.results))

	fmt.Printf("\n=== 执行摘要 ===\n")
	fmt.Printf("总任务数: %d\n", len(ts.tasks))
	fmt.Printf("成功: %d\n", successCount)
	fmt.Printf("失败: %d\n", failureCount)
	fmt.Printf("总耗时: %v\n", totalDuration)
	fmt.Printf("平均耗时: %v\n", avgDuration)
	fmt.Printf("成功率: %.1f%%\n", float64(successCount)/float64(len(ts.tasks))*100)
	fmt.Println()
}

// 获取执行结果
func (ts *TaskScheduler) GetResults() []TaskResult {
	return ts.results
}

// 重置调度器
func (ts *TaskScheduler) Reset() {
	ts.tasks = make([]Task, 0)
	ts.results = make([]TaskResult, 0)
}

// 示例任务函数
func createSampleTasks() []Task {
	return []Task{
		// 快速任务
		func() error {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("   快速任务完成")
			return nil
		},
		// 中等任务
		func() error {
			time.Sleep(300 * time.Millisecond)
			fmt.Println("   中等任务完成")
			return nil
		},
		// 慢速任务
		func() error {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("   慢速任务完成")
			return nil
		},
		// 可能失败的任务
		func() error {
			time.Sleep(200 * time.Millisecond)
			if time.Now().Unix()%2 == 0 {
				return fmt.Errorf("随机失败")
			}
			fmt.Println("   可能失败的任务完成")
			return nil
		},
		// 计算密集型任务
		func() error {
			start := time.Now()
			// 模拟计算
			for i := 0; i < 1000000; i++ {
				_ = i * i
			}
			fmt.Printf("   计算任务完成，耗时: %v\n", time.Since(start))
			return nil
		},
		// 网络请求模拟
		func() error {
			time.Sleep(400 * time.Millisecond)
			fmt.Println("   网络任务完成")
			return nil
		},
	}
}

// 创建会超时的任务
func createTimeoutTasks() []Task {
	return []Task{
		func() error {
			time.Sleep(2 * time.Second)
			fmt.Println("   这个任务应该会超时")
			return nil
		},
		func() error {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("   这个任务应该能完成")
			return nil
		},
	}
}

func main() {
	fmt.Println("=== Go 任务调度器演示 ===\n")

	// 演示1：串行 vs 并行执行
	demoSerialVsParallel()

	// 演示2：超时控制
	demoTimeoutControl()

	// 演示3：大量任务处理
	demoLargeTaskSet()
}

func demoSerialVsParallel() {
	fmt.Println("演示1：串行执行 vs 并行执行")

	tasks := createSampleTasks()

	// 串行执行
	scheduler1 := NewTaskScheduler(3)
	scheduler1.AddTasks(tasks)
	scheduler1.RunSerial()

	// 并行执行
	scheduler2 := NewTaskScheduler(3)
	scheduler2.AddTasks(tasks)
	scheduler2.RunParallel()
}

func demoTimeoutControl() {
	fmt.Println("演示2：任务超时控制")

	tasks := createTimeoutTasks()

	scheduler := NewTaskScheduler(2)
	scheduler.AddTasks(tasks)
	scheduler.SetTimeout(1 * time.Second) // 设置1秒超时
	scheduler.RunParallel()
}

func demoLargeTaskSet() {
	fmt.Println("演示3：大量任务处理")

	// 创建20个任务
	var tasks []Task
	for i := 0; i < 20; i++ {
		taskID := i
		tasks = append(tasks, func() error {
			// 随机睡眠时间 100-500ms
			sleepTime := time.Duration(100+(taskID*20)%400) * time.Millisecond
			time.Sleep(sleepTime)
			fmt.Printf("   任务 %d 完成\n", taskID)
			return nil
		})
	}

	scheduler := NewTaskScheduler(5) // 使用5个worker
	scheduler.AddTasks(tasks)
	scheduler.RunParallel()
}
