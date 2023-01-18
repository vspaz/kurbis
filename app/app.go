package app

import (
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/vspaz/simplelogger/pkg/logging"
	"kurbis/job"
	"kurbis/task"
	"kurbis/worker"
	"time"
)

func Run() {
	logger := logging.GetJsonLogger("info").Logger
	task_1 := task.Task{
		Id:     uuid.New(),
		Name:   "Task-1",
		State:  job.Pending,
		Image:  "Image-1",
		Memory: 1024,
		Disk:   1,
	}
	logger.Info(task_1)

	taskEvent := task.Event{
		Id:        uuid.New(),
		State:     job.Pending,
		Timestamp: time.Now(),
		Task:      task_1,
	}
	logger.Info(taskEvent)

	worker_1 := worker.Worker{
		Queue: *queue.New(),
	}
}
