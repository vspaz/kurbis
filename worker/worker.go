package worker

import (
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"kurbis/task"
	"time"
)

type Worker struct {
	Queue      queue.Queue
	UuidToTask map[uuid.UUID]task.Task
	TaskCount  int
	Logger     *logrus.Logger
}

func (w *Worker) GetStats() {
	w.Logger.Info("collecting stats")
}

func (w *Worker) RunTask() {
	w.Logger.Info("starting or stopping a task")
}

func (w *Worker) StartTask() {
	w.Logger.Info("starting a task")
}

func (w *Worker) StopTask(t task.Task) task.Result {
	config := task.NewConfig(&t)
	docker := task.NewDocker(&config)
	result := docker.Stop()
	if result.Error != nil {
		w.Logger.Errorf("failed to stop %v: %v", docker.Config.Runtime.ContainerId, result.Error)
	}
	t.FinishTime = time.Now().UTC()
	t.State = t.Completed
	w.UuidToTask[t.Id] = t
	w.Logger.Infof("stopped & removed container %v for task %v", t.ContainerId, t.ContainerId)
	return result
}
