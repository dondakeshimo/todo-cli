package json

import (
	ejson "encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/dondakeshimo/todo-cli/internal/values/remindtime"
	"github.com/dondakeshimo/todo-cli/internal/values/reminder"
)

// Task is a struct to write/read JSON.
type Task struct {
	ID         int    `json:"id"`
	Task       string `json:"task"`
	RemindTime string `json:"remindtime"`
	UUID       string `json:"uuid"`
	Reminder   string `json:"reminder"`
}

type Client struct {
	path string
	tasks []Task
}

const (
	jsonFile        = "todo/todo.json"
	defaultDataHome = ".local/share/"
)

func NewClient() (*Client, error) {
	c := new(Client)

	if err := c.exploreJSONPath(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) Read() ([]task.Task, error) {
	bytes, err := ioutil.ReadFile(c.path)
	if err != nil {
		return nil, err
	}

	var tjs []*Task
	if err:= ejson.Unmarshal(bytes, &tjs); err != nil {
		return nil, err
	}

	var ts []task.Task
	for _, t := range tjs {
		ts = append(ts, task.NewTask(
			t.ID, t.Task, remindtime.RemindTime(t.RemindTime), t.UUID, reminder.Reminder(t.Reminder)))
	}

	return ts, nil
}

// Write is a function that write Handler tasks to JSON file.
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

// exploreJSONPath is a function that set JSONPath.
func (c *Client) exploreJSONPath() error {
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

	c.path = jsonPath
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
		RemindTime: string(t.RemindTime()),
		UUID:       t.UUID(),
		Reminder:   string(t.Reminder()),
	}
}
