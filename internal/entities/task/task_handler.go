package task

import (
	"fmt"
	"sort"
)

// Repository is a interface to save tasks.
type Repository interface {
	Read() ([]Task, error)
	Write([]Task) error
}

// Handler is a struct that manage tasks.
type Handler struct {
	tasks      []Task
	repository Repository
}

// NewHandler is a constructor that make Handler.
func NewHandler(rep Repository) (*Handler, error) {
	t := new(Handler)

	t.repository = rep

	ts, err := t.repository.Read()
	if err != nil {
		return nil, err
	}

	t.tasks = ts

	return t, nil
}

// GetTask is a getter that get a task with id.
func (h *Handler) GetTask(id int) (Task, error) {
	if id > len(h.tasks) || id < 0 {
		return Task{}, fmt.Errorf("no exist id: %d", id)
	}
	return h.tasks[id-1], nil
}

// GetTasks is a getter that get all tasks.
func (h *Handler) GetTasks() []Task {
	return h.tasks
}

// UpdateTask is a function that overwrite task with id.
func (h *Handler) UpdateTask(id int, t Task) error {
	if id > len(h.tasks) || id < 0 {
		return fmt.Errorf("no exist id: %d", id)
	}

	h.tasks[id-1] = t
	return nil
}

// FindTaskWithUUID is a getter that get a task matched the given uuid.
func (h *Handler) FindTaskWithUUID(uuid string) (Task, error) {
	for _, t := range h.tasks {
		if uuid == t.uuid {
			return t, nil
		}
	}

	return Task{}, fmt.Errorf("no exist uuid: %s", uuid)
}

// AppendTask is a function that append task.
func (h *Handler) AppendTask(t Task) {
	h.tasks = append(h.tasks, t)
	h.align()
}

// RemoveTask is a function that remove a task matched the given uuid.
// Do not use this func in loop. Use RemoveTasks instead.
func (h *Handler) RemoveTask(id int) error {
	if id > len(h.tasks) || id <= 1 {
		return fmt.Errorf("invalid id [%d]", id)
	}

	h.tasks = append(h.tasks[:id-1], h.tasks[id:]...)
	h.align()

	return nil
}

// RemoveTasks is a function that remove tasks matched the given uuids.
func (h *Handler) RemoveTasks(ids []int) error {
	// below logic assume that ids is sorted ascending
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })

	// validate natural value
	if ids[0] < 1 {
		return fmt.Errorf("not natural value is invalid %v", ids)
	}

	// validate out of range, it's enough to check tail id
	if ids[len(ids)-1] > len(h.tasks) {
		return fmt.Errorf("no task with id %v", ids)
	}

	for i, id := range ids {
		h.tasks = append(h.tasks[:id-i-1], h.tasks[id-i:]...)
	}
	h.align()

	return nil
}

// Commit is a function that invoke repository.Write.
func (h *Handler) Commit() error {
	h.align()
	if err := h.repository.Write(h.tasks); err != nil {
		return err
	}

	return nil
}

// align is a function that redefine ids according to order.
func (h *Handler) align() {
	var ts []Task
	for i, t := range h.tasks {
		ts = append(ts, t.alterID(i+1))
	}
	h.tasks = ts
}
