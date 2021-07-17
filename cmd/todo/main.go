package main

import (
	"log"

	"github.com/dondakeshimo/todo-cli/pkg/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		log.Fatal(err)
	}
}
