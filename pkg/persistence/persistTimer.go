package persistence

import (
	"fmt"
	"log"
	"time"
)

type Task interface {
	Execute() error
}

type TimeScheduler struct {
	tasks  []Task
	ticker *time.Ticker
	quit   chan struct{}
}

func NewScheduler(tasks []Task, duration time.Duration) *TimeScheduler {
	return &TimeScheduler{
		tasks:  tasks,
		ticker: time.NewTicker(duration),
		quit:   make(chan struct{}),
	}
}

func (s *TimeScheduler) Run() {
	for {
		fmt.Println("定时任务启动....")
		select {
		case <-s.ticker.C:
			fmt.Println("开始执行任务")
			// 执行所有任务
			for _, task := range s.tasks {
				//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
				//defer cancel()

				err := task.Execute()
				if err != nil {
					log.Printf("Error executing task: %v\n", err)
				} else {
					log.Println("Task executed successfully")
				}
			}

		case <-s.quit:
			// 关闭定时任务
			s.ticker.Stop()
			return
		}
	}
}

func (s *TimeScheduler) Stop() {
	close(s.quit)
}
