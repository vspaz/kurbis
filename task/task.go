package task

import (
	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
	"kurbis/job"
	"time"
)

type Task struct {
	Id            uuid.UUID
	ContainerId   uuid.UUID
	Name          string
	State         job.State
	Image         string
	Memory        int
	Disk          int
	ExposedPorts  nat.PortSet
	PortBindings  map[string]string
	RestartPolicy string
	StartTime     time.Time
	FinishTime    time.Time
	Completed     job.State
}
