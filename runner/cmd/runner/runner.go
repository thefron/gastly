package main

import (
	"os/exec"

	"github.com/Buzzvil/gastly/runner/pkg/migration"
	"github.com/golang/glog"
)

type runner struct {
	GhostPath string
}

func (runner *runner) execGhost(migration *migration.Migration) (bool, error) {
	glog.Infof("mig_id=%d: Running %s", migration.ID, runner.GhostPath)
	cmd := exec.Command(runner.GhostPath)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		glog.Errorf("mig_id=%d: error getting stdout pipe for gh-ost exec (error: %s)", migration.ID, err)
		return false, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		glog.Errorf("mig_id=%d: error getting stderr pipe for gh-ost exec (error: %s)", migration.ID, err)
		return false, err
	}

	if err := cmd.Start(); err != nil {
		glog.Errorf("exec error: %s", err)
		return false, err
	}

	ghostLogChan := make(chan string)

	stdoutErrChan := make(chan error)
	stderrErrChan := make(chan error)

	go migration.WatchMigrationStdout(stdout, stdoutErrChan, ghostLogChan)
	go migration.WatchMigrationStderr(stderr, stderrErrChan, ghostLogChan)

	return true, nil
}

func (runner *runner) Start() error {
	return nil
}
