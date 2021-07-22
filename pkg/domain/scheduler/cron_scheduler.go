package scheduler

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	cronFileTmpPath = "/tmp/todo-cron.txt"
	cronPreSettings = `#!/bin/sh
# this cron is made by todo-cli
SHELL=/bin/bash
PATH=/sbin:/bin:/usr/sbin:/usr/bin

`
	cronCmdTemplate     = "{{minute}} {{hour}} {{day}} {{month}} * {{command}}\n"
	cronPreLineNum      = 4
	cronCmdStartlineNum = 6
)

// CronScheduler is a scheduler for command with cron.
type CronScheduler struct {
	templateVar map[string]string
	cronFile    string
}

// NewCronScheduler is a constructor that make new CronScheduler.
func NewCronScheduler() *CronScheduler {
	return &CronScheduler{}
}

// Register set schedule to cron.
func (cs *CronScheduler) Register(r Request) error {
	cs.cronFile = cronPreSettings + cronCmdTemplate
	cs.renderCmdTemplate(r)

	cron, err := getExistingCron()
	if err != nil {
		return err
	}

	cmds := cs.extractCronCmd(cron)
	for _, cmd := range cmds {
		cs.cronFile += cmd + "\n"
	}

	if err := cs.setCrontab(); err != nil {
		return err
	}

	return nil
}

// ClearExpired remove enough old cron line.
func (cs *CronScheduler) ClearExpired() {
	cs.cronFile = cronPreSettings

	// HACK: ignore error because this is post process
	var cron, _ = getExistingCron()

	cmds := cs.extractCronCmd(cron)
	for _, cmd := range cmds {
		cs.cronFile += cmd + "\n"
	}

	var _ = cs.setCrontab()
}

// RemoveWithID is a function that remove cron line with ID.
func (cs *CronScheduler) RemoveWithID(uuid string) error {
	cs.cronFile = cronPreSettings

	cron, err := getExistingCron()
	if err != nil {
		return err
	}

	cmds := cs.extractCronCmd(cron)
	cmds = filterWithUUID(cmds, uuid)
	for _, cmd := range cmds {
		cs.cronFile += cmd + "\n"
	}

	if err := cs.setCrontab(); err != nil {
		return err
	}

	return nil
}

// setCrontab set a cron file made by todo-cli to crontab.
func (cs *CronScheduler) setCrontab() error {
	if err := ioutil.WriteFile(cronFileTmpPath, []byte(cs.cronFile), 0644); err != nil {
		return err
	}

	if err := exec.Command("crontab", cronFileTmpPath).Run(); err != nil {
		return err
	}

	if err := os.Remove(cronFileTmpPath); err != nil {
		return err
	}

	return nil
}

// renderCmdTemplate render request information to cronCmdTemplate.
func (cs *CronScheduler) renderCmdTemplate(r Request) {
	cs.templateVar = make(map[string]string)

	cs.templateVar["{{command}}"] = r.Command

	cs.templateVar["{{month}}"] = strconv.Itoa(int(r.DateTime.Month()))
	cs.templateVar["{{day}}"] = strconv.Itoa(r.DateTime.Day())
	cs.templateVar["{{hour}}"] = strconv.Itoa(r.DateTime.Hour())
	cs.templateVar["{{minute}}"] = strconv.Itoa(r.DateTime.Minute())

	for k, v := range cs.templateVar {
		cs.cronFile = strings.Replace(cs.cronFile, k, v, 1)
	}
}

// extractCronCmd get command lines from existing crontab.
func (cs *CronScheduler) extractCronCmd(cron string) []string {
	cron = strings.TrimSpace(cron)
	cl := strings.Split(cron, "\n")
	if len(cl) < cronCmdStartlineNum {
		return []string{}
	}

	var ec []string
	for _, c := range cl[cronCmdStartlineNum-1:] {
		if strings.TrimSpace(c) == "" {
			continue
		}
		if isExpired(extractCronTime(c)) {
			continue
		}
		ec = append(ec, c)
	}

	return ec
}

// getExistingCron returns `crontab -l` result.
func getExistingCron() (string, error) {
	out, err := exec.Command("crontab", "-l").CombinedOutput()
	cron := string(out)
	if err != nil {
		// HACK: this is a ubuntu cron error message.
		//       defferent environment probably say the other error message.
		if strings.TrimSpace(cron) != fmt.Sprintf("no crontab for %s", os.Getenv("USER")) {
			return "", err
		}
	} else if !equalToPreSettings(cron) {
		return "", fmt.Errorf("user cron has already been installed and cannot overwrite your cron")
	}

	return cron, nil
}

// equalToPreSettings compares a existing cron file to cronPreSettings for validation to overwrite.
func equalToPreSettings(cron string) bool {
	user := strings.Split(cron, "\n")
	this := strings.Split(cronPreSettings, "\n")

	if len(user) < cronPreLineNum {
		return false
	}

	for i := 0; i < cronPreLineNum; i++ {
		if user[i] != this[i] {
			return false
		}
	}

	return true
}

// filterWithUUID remove a cron command line with the given UUID.
func filterWithUUID(cmds []string, uuid string) []string {
	var fc []string
	for _, c := range cmds {
		if strings.TrimSpace(c) == "" {
			continue
		}
		if strings.Contains(c, uuid) {
			continue
		}
		fc = append(fc, c)
	}

	return fc
}

// extractCronTime returns time.Time read from cron.
func extractCronTime(line string) time.Time {
	cl := strings.Split(line, " ")[:4]
	var mo, _ = strconv.Atoi(cl[3])
	var d, _ = strconv.Atoi(cl[2])
	var h, _ = strconv.Atoi(cl[1])
	var m, _ = strconv.Atoi(cl[0])
	return time.Date(time.Now().Year(), time.Month(mo), d, h, m, 0, 0, time.Local)
}
