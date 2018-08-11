package main

import (
	"flag"
	"os/exec"

	"github.com/golang/glog"

	"github.com/Buzzvil/gastly/runner/pkg/migration"
)

type runner struct {
	GhostPath string
}

func (runner *runner) execGhost(migration *migration.Migration) (bool, error) {
	glog.Infof("mig_id=%d: Running %s", migration.Id, runner.GhostPath)
	cmd := exec.Command(runner.GhostPath)
	if err := cmd.Start(); err != nil {
		glog.Errorf("exec error: %s", err)
		return false, err
	}
	return true, nil
}

func main() {
	runner := &runner{}
	flag.StringVar(&runner.GhostPath, "ghost-path", "/bin/gh-ost", "Path to gh-ost binary")
	flag.Parse()

	runner.execGhost(&migration.Migration{})
}
