package scheduler

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	plistPrefix = "com.dondakeshimo.todo-cli."
	plistExt    = ".plist"
	plistDir    = "Library/LaunchAgents/"
)

// LaunchdScheduler is a struct that is schedule command with launchd.
type LaunchdScheduler struct {
	templateVar map[string]string
	plist       string
}

// NewLaunchdScheduler is a constructor that make new LaunchdScheduler.
func NewLaunchdScheduler() *LaunchdScheduler {
	ls := new(LaunchdScheduler)
	ls.plist =
		`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>{{label}}</string>
    <key>ProgramArguments</key>
    <array>
{{command}}
    </array>
    <key>StartCalendarInterval</key>
    <dict>
        <key>Month</key>
        <integer>{{month}}</integer>
        <key>Day</key>
        <integer>{{day}}</integer>
        <key>Hour</key>
        <integer>{{hour}}</integer>
        <key>Minute</key>
        <integer>{{minute}}</integer>
    </dict>
    <key>StandardOutPath</key>
    <string>/tmp/todo.out</string>
    <key>StandardErrorPath</key>
    <string>/tmp/todo.err</string>
    <key>UserName</key>
    <string>{{user}}</string>
</dict>
</plist>`

	return ls
}

// Register is a function that set schedule to launchd.
func (ls *LaunchdScheduler) Register(r Request) error {
	// TODO: if datetime is over 1 year later, return error
	ls.buildPlist(r)

	plistName := plistPrefix + r.ID + "_" + strconv.FormatInt(r.DateTime.Unix(), 10) + plistExt
	var homeDir, _ = os.UserHomeDir()
	plistPath := filepath.Join(homeDir, plistDir, plistName)

	if err := ioutil.WriteFile(plistPath, []byte(ls.plist), 0644); err != nil {
		return err
	}

	if err := exec.Command("launchctl", "load", plistPath).Run(); err != nil {
		return err
	}

	return nil
}

// buildPlist is a function that fill plist template.
func (ls *LaunchdScheduler) buildPlist(r Request) {
	ls.templateVar = make(map[string]string)

	ls.templateVar["{{label}}"] = strconv.FormatInt(r.DateTime.Unix(), 10)

	ls.templateVar["{{command}}"] = buildPlistCommand(r.Command)

	ls.templateVar["{{month}}"] = strconv.Itoa(int(r.DateTime.Month()))
	ls.templateVar["{{day}}"] = strconv.Itoa(r.DateTime.Day())
	ls.templateVar["{{hour}}"] = strconv.Itoa(r.DateTime.Hour())
	ls.templateVar["{{minute}}"] = strconv.Itoa(r.DateTime.Minute())

	ls.templateVar["{{user}}"] = os.Getenv("USER")

	for k, v := range ls.templateVar {
		ls.plist = strings.Replace(ls.plist, k, v, 1)
	}
}

// buildCommand is a function that make command to set launchd.
func buildPlistCommand(str string) string {
	command := ""
	for _, s := range strings.Split(str, " ") {
		command = command + "        <string>" + s + "</string>\n"
	}
	return command
}

// ClearExpired is a function that remove plist file enough old.
func (ls *LaunchdScheduler) ClearExpired() {
	var homeDir, _ = os.UserHomeDir()
	plistPaths := filepath.Join(homeDir, plistDir, plistPrefix+"*"+plistExt)
	files, _ := filepath.Glob(plistPaths)

	for _, f := range files {
		_, t := extractIDAndTime(f)
		if !isExpired(t) {
			continue
		}

		if err := os.Remove(f); err != nil {
			continue
		}
	}
}

// RemoveWithID is a function that remove plist file with ID.
func (ls *LaunchdScheduler) RemoveWithID(id string) error {
	var homeDir, _ = os.UserHomeDir()
	plistPaths := filepath.Join(homeDir, plistDir, plistPrefix+"*"+plistExt)
	files, _ := filepath.Glob(plistPaths)

	for _, f := range files {
		d, _ := extractIDAndTime(f)
		if d == id {
			if err := os.Remove(f); err != nil {
				return err
			}

			return nil
		}
	}

	return fmt.Errorf("not found scheduler with id: %s", id)
}

// extractIDAndTime is a function that extract ID and time from plist file path.
func extractIDAndTime(path string) (string, time.Time) {
	f := filepath.Base(path)
	f = strings.Replace(f, plistPrefix, "", -1)
	f = strings.Replace(f, plistExt, "", -1)
	s := strings.Split(f, "_")

	t, _ := strconv.ParseInt(s[1], 10, 64)
	return s[0], time.Unix(t, 0)
}
