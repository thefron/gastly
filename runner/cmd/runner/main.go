package main

import (
	"flag"

	"github.com/Buzzvil/gastly/runner/pkg/migration"
)

func main() {
	runner := &runner{}
	flag.StringVar(&runner.GhostPath, "ghost-path", "/bin/gh-ost", "Path to gh-ost binary")
	flag.Parse()

	runner.execGhost(&migration.Migration{})
}
