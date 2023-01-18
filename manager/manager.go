package manager

import (
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"kurbis/task"
)

type Manager struct {
	Pending       queue.Queue
	NameToTasks   map[string][]task.Task
	NameToEvents  map[string][]task.Task
	Workers       []string
	WorkerToTasks map[string][]uuid.UUID
	TaskToWorker  map[uuid.UUID]string
}
