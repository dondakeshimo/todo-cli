package task

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const (
	jsonFile        = "todo/todo.json"
	defaultDataHome = ".local/share/"
)

// Handler is a struct that manage tasks.
type Handler struct {
	JSONPath string
	tasks    []*Task
}

// NewHandler is a constructor that make Handler.
func NewHandler() (*Handler, error) {
	t := new(Handler)

	if err := t.exploreJSONPath(); err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(t.JSONPath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, &t.tasks); err != nil {
		return nil, err
	}

	return t, nil
}

// GetTask is a getter that get a task with id.
func (h *Handler) GetTask(id int) *Task {
	if id > len(h.tasks) {
		return nil
	}
	return h.tasks[id-1]
}

// GetTasks is a getter that get all tasks.
func (h *Handler) GetTasks() []*Task {
	return h.tasks
}

// FindTaskWithUUID is a getter that get a task matched the given uuid.
func (h *Handler) FindTaskWithUUID(uuid string) *Task {
	for _, t := range h.tasks {
		if uuid == t.UUID {
			return t
		}
	}

	return nil
}

// AppendTask is a function that append task.
func (h *Handler) AppendTask(t *Task) {
	h.tasks = append(h.tasks, t)
	h.align()
}

// exploreJSONPath is a function that set JSONPath.
func (h *Handler) exploreJSONPath() error {
	dataHome := os.Getenv("XDG_DATA_HOME")
	var jsonPath string
	var homeDir, _ = os.UserHomeDir()
	if dataHome != "" {
		jsonPath = filepath.Join(dataHome, jsonFile)
	} else {
		jsonPath = filepath.Join(homeDir, defaultDataHome, jsonFile)
	}

	if err := createJSONFile(jsonPath); err != nil {
		return err
	}

	h.JSONPath = jsonPath
	return nil
}

// createJSONFile is a function that make new JSON file.
func createJSONFile(path string) error {
	if _, err := os.Stat(filepath.Dir(path)); err != nil {
		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return err
		}
	}

	if _, err := os.Stat(path); err != nil {
		if err := writeInitialSample(path); err != nil {
			return err
		}
	}

	return nil
}

// writeInitialSample is a function that make initial data to write JSON file.
func writeInitialSample(path string) error {
	uu, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	tasks := &[]*Task{
		{
			ID:         1,
			Task:       "deleting or modifying this task is your first TODO",
			RemindTime: "2099/01/01 00:00",
			UUID:       uu.String(),
		},
	}

	bytes, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, bytes, 0644); err != nil {
		return err
	}

	return nil
}

// Write is a function that write Handler tasks to JSON file.
func (h *Handler) Write() error {
	bytes, err := json.Marshal(&h.tasks)
	if err != nil {
		return nil
	}

	if err := ioutil.WriteFile(h.JSONPath, bytes, 0644); err != nil {
		return err
	}

	return nil
}

// RemoveTask is a function that remove a task matched the given uuid.
// Do not use this func in loop. Use RemoveTasks instead.
func (h *Handler) RemoveTask(id int) {
	if id > len(h.tasks) {
		return
	}

	h.tasks = append(h.tasks[:id-1], h.tasks[id:]...)
	h.align()
}

// RemoveTasks is a function that remove tasks matched the given uuids.
func (h *Handler) RemoveTasks(ids []int) {
	for i, id := range ids {
		if id-i > len(h.tasks) {
			continue
		}
		h.tasks = append(h.tasks[:id-i-1], h.tasks[id-i:]...)
	}
	h.align()
}

// align is a function that sort tasks.
func (h *Handler) align() {
	for i, t := range h.tasks {
		t.ID = i + 1
	}
}
