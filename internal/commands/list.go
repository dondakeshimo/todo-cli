package commands

import (
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/dondakeshimo/todo-cli/internal/entities/task"
	"github.com/urfave/cli/v2"
)

func List(c *cli.Context) error {
	th, err := task.NewTaskHandler()
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)

	tasks := [][]byte{}
	for _, task := range th.TaskList.Tasks {
		tasks = append(tasks, []byte(strconv.Itoa(task.Id)+"\t"+task.Task+"\t"+task.Deadline+"\n"))
	}

	defer w.Flush()

	w.Write([]byte("ID\tTask\tDeadline\n"))
	for _, t := range tasks {
		w.Write(t)
	}

	return nil
}
