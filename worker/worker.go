package worker

import (
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"kurbis/task"
)

type Worker struct {
	Queue      queue.Queue
	uuidToTask map[uuid.UUID]task.Task
	taskCount  int
	logger     *logrus.Logger
}

func (w *Worker) GetStats() {
	w.logger.Info("collecting stats")
}
