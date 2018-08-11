package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"

	"github.com/Buzzvil/gastly/runner/pkg/migration"
)

type runner struct {
	GhostPath string
}

func (runner *runner) execGhost(migration *migration.Migration) (bool, error) {
	cmd := exec.Command(runner.GhostPath)
	if err := cmd.Start(); err != nil {
		log.Println(fmt.Errorf("exec error: %s", err))
		return false, err
	}
	return true, nil
}

func main() {
	runner := new(runner)
	flag.StringVar(&runner.GhostPath, "ghost-path", "/gh-ost", "Path to gh-ost binary")

	runner.execGhost(&migration.Migration{})
}
