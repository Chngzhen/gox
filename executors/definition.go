package executors

// TaskHandleResult 任务处理结果类型。
type TaskHandleResult struct {
	task interface{}
	code uint8
	desc string
}

// TaskHandler 任务处理器类型。
type TaskHandler func(task interface{}) *TaskHandleResult
