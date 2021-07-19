// +build scenario

package scenario_test

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
	cleanup, err := setup()
	if err != nil {
		log.Fatalf("could not set xdg for test: %s", err.Error())
	}
	defer cleanup()

	m.Run()

}

func setup() (func(), error) {
	const binaryName = "todo"
	const binaryDir = "../../"
	const varName = "XDG_DATA_HOME"
	const jsonFilename = "todo.json"
	const configFilename = "config.yaml"
	const appDir = "todo"

	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	p := filepath.Join(dir, binaryDir, binaryName)
	_, err = os.Stat(p)
	if err != nil {
		log.Fatalf("could not find target binary [%s]: %s", p, err.Error())
	}

	xdg := os.Getenv(varName)

	if err := os.Setenv(varName, dir); err != nil {
		return nil, err
	}

	p = filepath.Join(dir, appDir, jsonFilename)
	_, err = os.Stat(p)
	if err == nil || os.IsExist(err) {
		if err := os.Remove(p); err != nil {
			log.Fatalf("could not remove [%s]: %s", p, err.Error())
		}
	}

	p = filepath.Join(dir, appDir, configFilename)
	_, err = os.Stat(p)
	if err == nil || os.IsExist(err) {
		if err := os.Remove(p); err != nil {
			log.Fatalf("could not remove [%s]: %s", p, err.Error())
		}
	}

	return func() {
		p := filepath.Join(dir, appDir, jsonFilename)
		_, err = os.Stat(p)
		if err == nil || os.IsExist(err) {
			if err := os.Remove(p); err != nil {
				log.Printf("could not remove [%s]: %s", p, err.Error())
			}
		}

		p = filepath.Join(dir, appDir, configFilename)
		_, err = os.Stat(p)
		if err == nil || os.IsExist(err) {
			if err := os.Remove(p); err != nil {
				log.Fatalf("could not remove [%s]: %s", p, err.Error())
			}
		}

		if err := os.Setenv(varName, xdg); err != nil {
			log.Printf("could not set original xdg, please check your xdg home.: %s", err.Error())
		}
	}, nil
}
