package main

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
    t.setJsonPath()

    bytes, err := ioutil.ReadFile(t.JsonPath)
    if err != nil {
        return nil, err
    }

    if err := json.Unmarshal(bytes, &t.TaskList); err != nil {
        return nil, err
    }

    return t, nil
}

func (t *TaskHandler) setJsonPath() {
    dataHome := os.Getenv("XDG_DATA_HOME")
    var jsonPath string
    var homeDir, _ = os.UserHomeDir()
    if dataHome != "" {
        jsonPath = filepath.Join(dataHome, jsonFile)
    } else {
        jsonPath = filepath.Join(homeDir, ".local/share/", jsonFile)
    }

    t.JsonPath = jsonPath
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
