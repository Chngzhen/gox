package executors

import (
	"fmt"
	"sync"
	"testing"
)

func TestGoroutinePool_Start(t *testing.T) {
	var waitGroup sync.WaitGroup

	// 任务数据队列
	taskQueue := make(chan interface{}, 10)
	// 任务处理器
	taskHandler := func(task interface{}) *TaskHandleResult {
		data := task.(int)
		return &TaskHandleResult{task: data, code: 0, desc: "成功"}
	}
	// 创建协程池
	goroutinePool := New(3, taskQueue, taskHandler, &waitGroup)

	// 启动协程池
	resultQueue := make(chan *TaskHandleResult, 10)
	goroutinePool.StartWithResult(resultQueue)

	// 添加任务。需要主动关闭信道，否则GoroutinePool的工作协程会死锁造成程序崩溃。
	for i := 0; i < 10; i++ {
		taskQueue <- i
	}
	close(taskQueue)

	waitGroup.Wait()

	// 读取结果。GoroutinePool会在所有工作协程结束后自动关闭结果队列，避免程序因死锁而崩溃。
	close(resultQueue)
	for result := range resultQueue {
		fmt.Printf("%+v\n", result)
	}
}
