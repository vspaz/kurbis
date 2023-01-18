package manager

import (
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"kurbis/task"
)

type Manager struct {
	Pending       queue.Queue
	NameToTasks   map[string][]task.Task
	NameToEvents  map[string][]task.Task
	Workers       []string
	WorkerToTasks map[string][]uuid.UUID
	TaskToWorker  map[uuid.UUID]string
	Logger        *logrus.Logger
}

func (m *Manager) SelectWorker() {
	m.Logger.Info("selecting a worker")
}

func (m *Manager) UpdateTasks() {
	m.Logger.Info("updating tasks")
}

func (m *Manager) SendWork() {
	m.Logger.Info("sending workers")
}
