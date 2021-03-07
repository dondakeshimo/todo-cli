package main

import (
	"log"
	"os"

	"github.com/dondakeshimo/todo-cli/internal/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	flagTask := &cli.StringFlag{
		Name:    "task",
		Aliases: []string{"t"},
		Usage:   "task contents",
	}
	flagRemindTime := &cli.StringFlag{
		Name:    "remind_time",
		Aliases: []string{"d"},
		Usage:   "remind_time (2021/3/3 03:03, 2021/3/3, +2h3m, task-4h15m)",
	}
	flagReminder := &cli.StringFlag{
		Name:    "reminder",
		Aliases: []string{"r"},
		Usage:   "choose reminder from [macos]",
	}
	flagID := &cli.IntFlag{
		Name:     "id",
		Aliases:  []string{"i"},
		Usage:    "task's ID",
		Required: true,
	}
	flagIDs := &cli.IntSliceFlag{
		Name:     "ids",
		Aliases:  []string{"i"},
		Usage:    "task's IDs  (ex. $ todo c -i 2 -i 3 -i 5 )",
		Required: true,
	}
	flagUUID := &cli.StringFlag{
		Name:     "uuid",
		Usage:    "task's UUID",
		Required: true,
	}
	flagRemoveReminder := &cli.BoolFlag{
		Name:  "remove-reminder",
		Usage: "remove reminder. this option overrides reminder option",
	}

	app := &cli.App{
		Name:                 "todo",
		Usage:                "Manage Your TODO",
		Version:              "0.0.1",
		EnableBashCompletion: true,
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
					flagIDs,
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
					flagRemoveReminder,
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
