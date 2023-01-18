package task

import (
	"github.com/google/uuid"
	"kurbis/job"
)

type Task struct {
	Id    uuid.UUID
	Name  string
	State job.State
}
