package app

import (
	"github.com/google/uuid"
	"github.com/vspaz/simplelogger/pkg/logging"
	"kurbis/job"
	"kurbis/task"
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
}
