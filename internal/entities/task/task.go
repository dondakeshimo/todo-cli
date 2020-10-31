package task

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	jsonFile = "todo/todo.json"
)

type Task struct {
	ID       int    `json:"id"`
	Task     string `json:"task"`
	Deadline string `json:"deadline"`
}

type List struct {
	Tasks []*Task `json:"tasklist"`
}

type Handler struct {
	JSONPath string
	TaskList *List
}

func NewHandler() (*Handler, error) {
	t := new(Handler)

	if err := t.exploreJSONPath(); err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(t.JSONPath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, &t.TaskList); err != nil {
		return nil, err
	}

	return t, nil
}

func (h *Handler) exploreJSONPath() error {
	dataHome := os.Getenv("XDG_DATA_HOME")
	var jsonPath string
	var homeDir, _ = os.UserHomeDir()
	if dataHome != "" {
		jsonPath = filepath.Join(dataHome, jsonFile)
	} else {
		jsonPath = filepath.Join(homeDir, ".local/share/", jsonFile)
	}

	if err := createJSONFile(jsonPath); err != nil {
		return err
	}

	h.JSONPath = jsonPath
	return nil
}

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

func writeInitialSample(path string) error {
	taskList := &List{[]*Task{
		{
			ID:       1,
			Task:     "deleting or modifying this task is your first TODO",
			Deadline: "2099/01/01 00:00",
		},
	}}

	bytes, err := json.Marshal(taskList)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, bytes, 0644); err != nil {
		return err
	}

	return nil
}

func (h *Handler) Write() error {
	bytes, err := json.Marshal(&h.TaskList)
	if err != nil {
		return nil
	}

	if err := ioutil.WriteFile(h.JSONPath, bytes, 0644); err != nil {
		return err
	}

	return nil
}

func (h *Handler) Remove(id int) {
	if id > len(h.TaskList.Tasks) {
		return
	}

	h.TaskList.Tasks = append(h.TaskList.Tasks[:id], h.TaskList.Tasks[id+1:]...)
	h.align()
}

func (h *Handler) align() {
	for i, t := range h.TaskList.Tasks {
		t.ID = i + 1
	}
}
