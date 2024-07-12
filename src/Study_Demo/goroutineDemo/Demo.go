package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 任务类型
type Task struct {
	ID   int
	Data string
}

// Result 结果类型
type Result struct {
	TaskID int
	Status string
}

// Heartbeat 心跳消息类型
type Heartbeat struct {
	WorkerID int
}

func main() {
	// 无缓冲的任务分配 channel
	taskChan := make(chan Task)
	// 缓冲为 5 的结果收集 channel
	resultChan := make(chan Result, 5)
	// 心跳接收 channel
	heartbeatChan := make(chan Heartbeat)
	// 未完成任务存储 channel
	pendingTaskChan := make(chan Task)

	var wg sync.WaitGroup

	// 模拟 3 个工作节点
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, taskChan, resultChan, heartbeatChan, pendingTaskChan, &wg)
	}

	// 主节点分配任务
	go master(taskChan, pendingTaskChan)

	// 主节点处理心跳和故障
	go handleHeartbeatsAndFailures(heartbeatChan, taskChan, pendingTaskChan)

	wg.Wait()
}

func master(taskChan chan<- Task, pendingTaskChan chan<- Task) {
	for i := 1; i <= 10; i++ {
		task := Task{ID: i, Data: fmt.Sprintf("Task %d", i)}
		taskChan <- task
		time.Sleep(100 * time.Millisecond)
	}
	close(taskChan)
}

func worker(workerID int, taskChan <-chan Task, resultChan chan<- Result, heartbeatChan chan<- Heartbeat, pendingTaskChan chan<- Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case task, ok := <-taskChan:
			if ok {
				fmt.Printf("Worker %d received task: %v\n", workerID, task)
				// 模拟处理任务
				time.Sleep(500 * time.Millisecond)
				result := Result{TaskID: task.ID, Status: "Done"}
				resultChan <- result
			} else {
				return
			}
		default:
			// 发送心跳
			heartbeatChan <- Heartbeat{WorkerID: workerID}
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func handleHeartbeatsAndFailures(heartbeatChan <-chan Heartbeat, taskChan chan<- Task, pendingTaskChan chan<- Task) {
	// 记录每个工作节点最后一次心跳的时间
	lastHeartbeatTimes := make(map[int]time.Time)

	for {
		select {
		case hb := <-heartbeatChan:
			lastHeartbeatTimes[hb.WorkerID] = time.Now()
		default:
			// 检查工作节点是否故障
			for workerID, lastTime := range lastHeartbeatTimes {
				if time.Since(lastTime) > 500*time.Millisecond {
					fmt.Printf("Worker %d is considered failed\n", workerID)
					// 将该工作节点未完成的任务重新放入待分配任务通道
					reassignPendingTasks(workerID, pendingTaskChan)
					delete(lastHeartbeatTimes, workerID)
				}
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func reassignPendingTasks(workerID int, pendingTaskChan chan<- Task) {
	// 模拟获取该故障工作节点未完成的任务，并重新放入待分配通道
	fmt.Printf("Reassigning tasks of failed worker %d\n", workerID)
}
