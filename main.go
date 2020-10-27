package main

import (
    "os"

    "github.com/urfave/cli"
)

func main() {
    app := cli.NewApp()

    app.Name = "todo"
    app.Usage = "Manage your TODO"
    app.Version = "0.0.1"

    app.Run(os.Args)
}
