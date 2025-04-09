package queue

import (
	"sync"
	"time"
)

type TaskManager struct { 
	mutex sync.Mutex
	tasks []Task
	readyChan chan Task
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks: []Task{},
		readyChan: make(chan Task, 100),
	}
}

func (tm_this *TaskManager) AddTask(t Task) {
	tm_this.mutex.Lock()

	// even if a panic or return happens here, unlock will run
	defer tm_this.mutex.Unlock()
	tm_this.tasks = append(tm_this.tasks, t)
}

func (tm_this *TaskManager) StartScheduler() {
	go func()  {
		
		// infinite loop the keep checking if new entry is there or not
		for {
			tm_this.mutex.Lock()
			now := time.Now()
			var remaining []Task

			for _, t := range tm_this.tasks {
				if now.After(t.AvailableAt) || now.Equal(t.AvailableAt) {
					tm_this.readyChan <- t
				} else {
					remaining = append(remaining, t)
				}
			}
			
			tm_this.tasks = remaining
			tm_this.mutex.Unlock()
			time.Sleep(2 * time.Second)

		}
	}()
}

func (tm_this *TaskManager) GetReadyTask()(Task, bool) {
	// select is a special keyword to handle channels
	select {
	case t:= <-tm_this.readyChan:
		return t, true
	default:
			return Task{}, false
	}
}