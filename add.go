package main

import (
    "os"
    "fmt"
    "time"

    "github.com/urfave/cli/v2"
)

func Add(c *cli.Context) error {
    th, err := NewTaskHandler()
    if err != nil {
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
