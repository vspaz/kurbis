package worker

import (
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"kurbis/task"
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

}
