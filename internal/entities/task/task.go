package task

import (
    "os"
    "encoding/json"
    "io/ioutil"
    "path/filepath"
)

const (
    jsonFile = "todo/todo.json"
)

type Task struct {
    Id int `json:"id"`
    Task string `json:"task"`
    Deadline string `json:"deadline"`
}

type TaskList struct {
    Tasks []*Task `json:"tasklist"`
}

type TaskHandler struct {
    JsonPath string
    TaskList *TaskList
}

func NewTaskHandler() (*TaskHandler, error) {
    t := new(TaskHandler)

    if err := t.exploreJsonPath(); err != nil {
        return nil, err
    }

    bytes, err := ioutil.ReadFile(t.JsonPath)
    if err != nil {
        return nil, err
    }

    if err := json.Unmarshal(bytes, &t.TaskList); err != nil {
        return nil, err
    }

    return t, nil
}

func (t *TaskHandler) exploreJsonPath() error {
    dataHome := os.Getenv("XDG_DATA_HOME")
    var jsonPath string
    var homeDir, _ = os.UserHomeDir()
    if dataHome != "" {
        jsonPath = filepath.Join(dataHome, jsonFile)
    } else {
        jsonPath = filepath.Join(homeDir, ".local/share/", jsonFile)
    }

    if err := createJsonFile(jsonPath); err != nil {
        return err
    }

    t.JsonPath = jsonPath
    return nil
}

func createJsonFile(path string) error {
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
    taskList := &TaskList{[]*Task{
        {
            Id: 1,
            Task: "deleting or modifying this task is your first TODO",
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

func (t *TaskHandler) Write() error {
    bytes, err := json.Marshal(&t.TaskList)
    if err != nil {
        return nil
    }

    if err := ioutil.WriteFile(t.JsonPath, bytes, 0644); err != nil {
        return err
    }

    return nil
}
