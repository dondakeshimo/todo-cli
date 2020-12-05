package scenario_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
	cleanup, err := setup()
	if err != nil {
		fmt.Printf("could not set xdg for test: %s", err.Error())
	}
	defer cleanup()

	m.Run()

}

func setup() (func(), error) {
	const varName = "XDG_DATA_HOME"
	const jsonFilename = "todo.json"
	const jsonDir = "todo"

	xdg := os.Getenv(varName)

	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	if err := os.Setenv(varName, dir); err != nil {
		return nil, err
	}

	return func() {
		p := filepath.Join(dir, jsonDir, jsonFilename)
		_, err = os.Stat(p)
		if err == nil || os.IsExist(err) {
			fmt.Printf("remove %s", p)
			if err := os.Remove(p); err != nil {
				fmt.Printf("could not remove [%s]: %s", p, err.Error())
			}
		}

		if err := os.Setenv(varName, xdg); err != nil {
			fmt.Printf("could not set original xdg, please check your xdg home.: %s", err.Error())
		}
	}, nil
}
