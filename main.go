package main

import (
    "os"
    "log"
    "fmt"
    "encoding/json"
    "io/ioutil"
    "time"

    "github.com/urfave/cli/v2"
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

func (t *TaskHandler) Read() error {
    bytes, err := ioutil.ReadFile(t.JsonPath)
    if err != nil {
        return err
    }

    if err := json.Unmarshal(bytes, &t.TaskList); err != nil {
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

func List(c *cli.Context) error {
    var th = TaskHandler{JsonPath: "todo.json"}
    if err := th.Read(); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
    for _, task := range th.TaskList.Tasks {
        fmt.Printf("id: %d, task: %s, deadline: %s", task.Id, task.Task, task.Deadline)
    }
    return nil
}

func Add(c *cli.Context) error {
    var th = TaskHandler{JsonPath: "todo.json"}
    if err := th.Read(); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }

    // validation
    if c.String("task") == "" {
        fmt.Fprintln(os.Stderr, "task could not be empty")
    }


    if c.String("deadline") != "" {
        layout := "2006/01/02 15:04"
        _, err := time.Parse(layout, c.String("deadline"))
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
    }

    tl := th.TaskList
    tl.Tasks = append(tl.Tasks, &Task{
        Id: len(tl.Tasks) + 1,
        Task: c.String("task"),
        Deadline: c.String("deadline"),
    })

    if err := th.Write(); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }

    return nil
}

func main() {
    app := &cli.App{
        Name: "todo",
        Usage: "Manage Your TODO",
        Version: "0.0.1",
        Commands: []*cli.Command{
            {
                Name: "list",
                Aliases: []string{"l"},
                Usage: "list tasks",
                Action: List,
            },
            {
                Name: "add",
                Aliases: []string{"a"},
                Usage: "add a task",
                Action: Add,
                Flags: []cli.Flag{
                    &cli.StringFlag{
                        Name: "task",
                        Aliases: []string{"t"},
                        Usage: "write task contents",
                    },
                    &cli.StringFlag{
                        Name: "deadline",
                        Aliases: []string{"d"},
                        Usage: "write deadline format -> 2020/10/27 19:35",
                    },
                },
            },
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
