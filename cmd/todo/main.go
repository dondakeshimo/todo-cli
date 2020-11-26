package main

import (
	"log"
	"os"

	"github.com/dondakeshimo/todo-cli/internal/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	flagTask := &cli.StringFlag{
		Name:     "task",
		Aliases:  []string{"t"},
		Usage:    "write task contents",
		Required: true,
	}
	flagRemindTime := &cli.StringFlag{
		Name:    "remind_time",
		Aliases: []string{"d"},
		Usage:   "write remind_time format -> 2020/10/27 19:35",
	}
	flagReminder := &cli.StringFlag{
		Name:    "reminder",
		Aliases: []string{"d"},
		Usage:   "choose reminder from [macos, slack]",
	}
	flagID := &cli.StringFlag{
		Name:     "id",
		Aliases:  []string{"i"},
		Usage:    "task's ID",
		Required: true,
	}
	flagUUID := &cli.StringFlag{
		Name:     "uuid",
		Usage:    "task's UUID",
		Required: true,
	}

	app := &cli.App{
		Name:    "todo",
		Usage:   "Manage Your TODO",
		Version: "0.0.1",
		Commands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List tasks",
				Action:  commands.List,
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Add a task",
				Action:  commands.Add,
				Flags: []cli.Flag{
					flagTask,
					flagRemindTime,
					flagReminder,
				},
			},
			{
				Name:    "close",
				Aliases: []string{"c"},
				Usage:   "Close a task",
				Action:  commands.Close,
				Flags: []cli.Flag{
					flagID,
				},
			},
			{
				Name:    "modify",
				Aliases: []string{"m"},
				Usage:   "Modify a task",
				Action:  commands.Modify,
				Flags: []cli.Flag{
					flagID,
					flagTask,
					flagRemindTime,
					flagReminder,
				},
			},
			{
				Name:   "notify",
				Usage:  "Notify a task (basicaly be used by system)",
				Action: commands.Notify,
				Flags: []cli.Flag{
					flagUUID,
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
