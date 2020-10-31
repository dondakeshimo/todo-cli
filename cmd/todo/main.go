package main

import (
	"log"
	"os"

	"github.com/dondakeshimo/todo-cli/internal/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "todo",
		Usage:   "Manage Your TODO",
		Version: "0.0.1",
		Commands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list tasks",
				Action:  commands.List,
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task",
				Action:  commands.Add,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "task",
						Aliases: []string{"t"},
						Usage:   "write task contents",
					},
					&cli.StringFlag{
						Name:    "deadline",
						Aliases: []string{"d"},
						Usage:   "write deadline format -> 2020/10/27 19:35",
					},
				},
			},
			{
				Name:    "close",
				Aliases: []string{"c"},
				Usage:   "close a task",
				Action:  commands.Close,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "id",
						Aliases: []string{"i"},
						Usage:   "task's ID",
					},
				},
			},
			{
				Name:    "modify",
				Aliases: []string{"m"},
				Usage:   "modify a task",
				Action:  commands.Modify,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "id",
						Aliases: []string{"i"},
						Usage:   "task's ID",
					},
					&cli.StringFlag{
						Name:    "task",
						Aliases: []string{"t"},
						Usage:   "write task contents",
					},
					&cli.StringFlag{
						Name:    "deadline",
						Aliases: []string{"d"},
						Usage:   "write deadline format -> 2020/10/27 19:35",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
