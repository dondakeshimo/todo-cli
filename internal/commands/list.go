package commands

import (
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/urfave/cli/v2"
)

func List(c *cli.Context) error {
	h, err := task.NewHandler()
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)

	tasks := [][]byte{}
	for _, task := range h.TaskList.Tasks {
		tasks = append(tasks, []byte(strconv.Itoa(task.ID)+"\t"+task.Task+"\t"+task.Deadline+"\n"))
	}

	defer w.Flush()

	if _, err := w.Write([]byte("ID\tTask\tDeadline\n")); err != nil {
		return err
	}

	for _, t := range tasks {
		if _, err := w.Write(t); err != nil {
			return err
		}
	}

	return nil
}
