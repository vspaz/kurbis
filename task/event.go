package task

import (
	"github.com/google/uuid"
	"kurbis/job"
	"time"
)

type Event struct {
	Id        uuid.UUID
	State     job.State
	Timestamp time.Time
	Task      Task
}
