package migration

import (
	"bufio"
	"errors"
	"fmt"
	"io"

	"github.com/golang/glog"
)

// Migration wraps attributes required for running gh-ost migration
type Migration struct {
	ID             int
	Status         int
	Host           string
	Port           int
	Database       string
	Table          string
	AlterStatement string
}

// WatchMigrationStdout scans stdout of migraiton and logs to file
func (migration *Migration) WatchMigrationStdout(stdoutPipe io.Reader, errChan chan error, ghostLogChan chan string) {
	scanner := bufio.NewScanner(stdoutPipe)
	for scanner.Scan() {
		line := scanner.Text()
		ghostLogChan <- fmt.Sprintf("stdout: %s", line)
	}
	if err := scanner.Err(); err != nil {
		glog.Errorf("mig_id=%d: error getting stdout of pt-osc (error: %s)", migration.ID, err)
		errChan <- errors.New("migration: failed to get stdout of gh-ost")
		return
	}

	errChan <- nil
}

// WatchMigrationStderr scans stdout of migraiton and logs to file
func (migration *Migration) WatchMigrationStderr(stderrPipe io.Reader, errChan chan error, ghostLogChan chan string) {
	scanner := bufio.NewScanner(stderrPipe)
	var wasError bool

	for scanner.Scan() {
		line := scanner.Text()
		wasError = true
		ghostLogChan <- fmt.Sprintf("stderr: %s", line)
	}
	if err := scanner.Err(); err != nil {
		glog.Errorf("mig_id=%d: error getting stderr of gh-ost (error: %s)", migration.ID, err)
		errChan <- errors.New("migration: failed to get stderr of gh-ost")
		return
	}

	if wasError {
		glog.Errorf("mig_id=%d: stderr is not empty. Something went wrong", migration.ID)
		errChan <- errors.New("migration: gh-ost stderr is not as expected")
		return
	}

	errChan <- nil
}
