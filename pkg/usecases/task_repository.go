package usecases

import (
	"github.com/dondakeshimo/todo-cli/pkg/domain/task"
)

var taskRepository task.Repository

// SetRepository is a setter of repository.
func SetRepository(tr task.Repository) {
	taskRepository = tr
}
