package worker

import (
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"kurbis/task"
)

type Worker struct {
	Queue      queue.Queue
	uuidToTask map[uuid.UUID]task.Task
	taskCount  int
}
