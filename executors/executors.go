// Package xexecutors
// @Title 协程池
// @Description 支持多协程执行任务。
// @Author chngzhen
// @CreatedOn 2021-11-18
// @UpdatedOn 2022-06-23
package executors

import (
	"sync"
)

// GoroutinePool 协程池类型。
type GoroutinePool struct {
	workerNum   uint8              // 工作协程的数量。
	taskQueue   <-chan interface{} // 任务队列。
	taskHandler TaskHandler        // 任务处理器。
	waitGroup   *sync.WaitGroup    // 同步阻塞器。
}

// New 协程池的构造器。
//
// @param workerNum   uint8          工作协程的数量。
// @param queue chan  interface{}    任务队列。注意：任务生产者需要在生产结束后主动关闭任务队列，否则工作协程会死锁致程序崩溃。
// @param taskHandler TaskHandler    任务处理器。
// @param waitGroup   sync.WaitGroup 同步阻塞器。若调用者依赖 sync.WaitGroup 维持程序的运行，则需要将其传入；否则，可以忽略。
// @return 协程池实例。
func New(workerNum uint8, queue <-chan interface{}, taskHandler TaskHandler, waitGroup *sync.WaitGroup) *GoroutinePool {
	goroutinePool := &GoroutinePool{
		workerNum:   workerNum,
		taskQueue:   queue,
		taskHandler: taskHandler,
		waitGroup:   waitGroup,
	}
	return goroutinePool
}

// Start 启动协程池。协程池在构造结束后不会立即执行，需要调用本方法启动。
func (t *GoroutinePool) Start() {
	var i uint8 = 0
	for ; i < t.workerNum; i++ {
		if t.waitGroup != nil {
			t.waitGroup.Add(1)
		}

		go t.doExecute()
	}
}

// StartWithResult 启动协程池。协程池在构造结束后不会立即执行，需要调用本方法启动。
//
// @param results chan 任务处理结果队列。
func (t *GoroutinePool) StartWithResult(results chan<- *TaskHandleResult) {
	var i uint8 = 0
	for ; i < t.workerNum; i++ {
		if t.waitGroup != nil {
			t.waitGroup.Add(1)
		}

		go t.doExecuteWithResult(results)
	}
}

func (t *GoroutinePool) doExecute() {
	for task := range t.taskQueue {
		t.taskHandler(task)
	}

	if t.waitGroup != nil {
		t.waitGroup.Done()
	}
}

func (t *GoroutinePool) doExecuteWithResult(results chan<- *TaskHandleResult) {
	for task := range t.taskQueue {
		results <- t.taskHandler(task)
	}

	if t.waitGroup != nil {
		t.waitGroup.Done()
	}
}
