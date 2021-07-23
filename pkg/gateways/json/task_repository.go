package json

import (
	ejson "encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/dondakeshimo/todo-cli/pkg/domain/task"
	"github.com/google/uuid"
)

// Task is a struct to write/read JSON.
type Task struct {
	ID         int    `json:"id"`
	Task       string `json:"task"`
	Group      string `json:"group"`
	RemindTime string `json:"remindtime"`
	UUID       string `json:"uuid"`
	Reminder   string `json:"reminder"`
	Priority   int    `json:"priority"`
}

// Client is a struct to manage JSON I/O.
type Client struct {
	path string
}

// NewClient is a constructor for Client.
func NewClient(path string) (*Client, error) {
	if err := createJSONFileIfNotExist(path); err != nil {
		return nil, err
	}
	return &Client{path: path}, nil
}

// Read is a function that read tasks from JSON file.
func (c *Client) Read() ([]task.Task, error) {
	bytes, err := ioutil.ReadFile(c.path)
	if err != nil {
		return nil, err
	}

	var tjs []*Task
	if err := ejson.Unmarshal(bytes, &tjs); err != nil {
		return nil, err
	}

	var ts []task.Task
	for _, t := range tjs {
		ts = append(ts, task.NewTask(
			t.ID, t.Task, t.Group, task.RemindTime(t.RemindTime), t.UUID, task.Reminder(t.Reminder), t.Priority))
	}

	return ts, nil
}

// Write is a function that write tasks to JSON file.
func (c *Client) Write(ts []task.Task) error {
	var tjs []*Task
	for _, t := range ts {
		tj := convertToJSONTask(t)
		tjs = append(tjs, &tj)
	}

	bytes, err := ejson.Marshal(&tjs)
	if err != nil {
		return nil
	}

	if err := ioutil.WriteFile(c.path, bytes, 0644); err != nil {
		return err
	}

	return nil
}

// createJSONFileIfNotExist is a function that make new JSON file if not exist.
func createJSONFileIfNotExist(path string) error {
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
			RemindTime: "2099/1/1 00:00",
			UUID:       uu.String(),
		},
	}

	bytes, err := ejson.Marshal(tasks)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, bytes, 0644); err != nil {
		return err
	}

	return nil
}

func convertToJSONTask(t task.Task) Task {
	return Task{
		ID:         t.ID(),
		Task:       t.Task(),
		Group:      t.Group(),
		RemindTime: string(t.RemindTime()),
		UUID:       t.UUID(),
		Reminder:   string(t.Reminder()),
		Priority:   t.Priority(),
	}
}
