package scheduler

import (
	"io/ioutil"
	"os"
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

type LaunchdScheduler struct {
	templateVar map[string]string
	plist       string
}

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

func (ls *LaunchdScheduler) Register(r *Request) error {
	// TODO: if datetime is over 1 year later, return error
	ls.buildPlist(r)

	plistName := plistPrefix + strconv.FormatInt(r.DateTime.Unix(), 10) + plistExt
	var homeDir, _ = os.UserHomeDir()
	plistPath := filepath.Join(homeDir, plistDir, plistName)

	if err := ioutil.WriteFile(plistPath, []byte(ls.plist), 0644); err != nil {
		return err
	}

	return nil
}

func (ls *LaunchdScheduler) buildPlist(r *Request) {
	ls.templateVar["{{label}}"] = strconv.FormatInt(r.DateTime.Unix(), 10)

	ls.templateVar["{{command}}"] = buildCommand(r.Command)

	ls.templateVar["{{month}}"] = strconv.Itoa(int(r.DateTime.Month()))
	ls.templateVar["{{day}}"] = strconv.Itoa(r.DateTime.Day())
	ls.templateVar["{{hour}}"] = strconv.Itoa(r.DateTime.Hour())
	ls.templateVar["{{minute}}"] = strconv.Itoa(r.DateTime.Minute())

	ls.templateVar["{{user}}"] = os.Getenv("USER")

	for k, v := range ls.templateVar {
		ls.plist = strings.Replace(ls.plist, k, v, 1)
	}
}

func buildCommand(str string) string {
	command := ""
	for _, s := range strings.Split(str, " ") {
		command = command + "        <string>" + s + "</string>\n"
	}
	return command
}

func (ls *LaunchdScheduler) ClearExpired() {
	var homeDir, _ = os.UserHomeDir()
	plistPaths := filepath.Join(homeDir, plistDir, plistPrefix, "*", plistExt)
	files, _ := filepath.Glob(plistPaths)

	for _, f := range files {
		t := extractTimeFromPath(f)
		if !isExpired(t) {
			continue
		}
		os.Remove(f)
	}
}

func extractTimeFromPath(path string) time.Time {
	f := filepath.Base(path)
	f = strings.Replace(f, plistPrefix, "", -1)
	f = strings.Replace(f, plistExt, "", -1)

	t, _ := strconv.ParseInt(f, 10, 64)
	return time.Unix(t, 0)
}

func isExpired(t time.Time) bool {
	deadline := time.Now().Add(time.Duration(10) * time.Minute).Unix()
	return t.Unix() < deadline
}
