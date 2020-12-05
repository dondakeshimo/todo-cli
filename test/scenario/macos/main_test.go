package scenario_test

import (
	"fmt"
	"testing"
	"os"
)

func TestMain(m *testing.M) {
	cleanup, err := setEnv()
	if err != nil {
		fmt.Printf("could not set xdg for test: %s", err.Error())
	}
	defer cleanup()

	m.Run()

}

func setEnv() (func(), error) {
	const varName = "XDG_DATA_HOME"
	xdg := os.Getenv(varName)

	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	if err := os.Setenv(varName, dir); err != nil {
		return nil, err
	}

	return func () {
		if err := os.Setenv(varName, xdg); err != nil {
			fmt.Printf("could not set original xdg, please check your xdg home.: %s", err.Error())
		}
	}, nil
}
