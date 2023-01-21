package worker

import (
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"kurbis/job"
	"kurbis/task"
	"sync"
	"time"
)

var mtx sync.Mutex

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

func (w *Worker) StartTask(t task.Task) task.Result {
	config := task.NewConfig(&t)
	docker := task.NewDocker(&config)
	result := docker.Run()
	if result.Error != nil {
		w.Logger.Infof("failed to run task %v: %v", t.Id, result.Error)
		t.State = job.Failed
		mtx.Lock()
		w.UuidToTask[t.Id] = t
		mtx.Lock()
		return result
	}
	t.ContainerId = result.ContainerId
	t.State = job.Running
	mtx.Lock()
	w.UuidToTask[t.Id] = t
	mtx.Lock()
	return result
}

func (w *Worker) StopTask(t task.Task) task.Result {
	config := task.NewConfig(&t)
	docker := task.NewDocker(&config)
	result := docker.Stop()
	if result.Error != nil {
		w.Logger.Errorf("failed to stop %v: %v", t.Id, result.Error)
	}
	t.FinishTime = time.Now().UTC()
	t.State = job.Completed
	mtx.Lock()
	w.UuidToTask[t.Id] = t
	mtx.Unlock()
	w.Logger.Infof("stopped & removed container %v for task %v", t.ContainerId, t.ContainerId)
	return result
}
