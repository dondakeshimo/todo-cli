package commands

import (
    "os"
    "fmt"

    "github.com/dondakeshimo/todo-cli/internal/entities/task"
    "github.com/urfave/cli/v2"
)

func List(c *cli.Context) error {
    th, err := task.NewTaskHandler()
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
    for _, task := range th.TaskList.Tasks {
        fmt.Printf("id: %d, task: %s, deadline: %s", task.Id, task.Task, task.Deadline)
    }
    return nil
}
