package tasks

import (
	"fmt"
	"io"

	"github.com/gookit/color"
)

var (
	styleDone  = color.New(color.FgGreen, color.Bold).Render
	styleError = color.New(color.FgRed, color.Bold).Render
)

// The Task struct holds the necessary information to produce output for users that execute a multi-step process.
type Task struct {
	ErrorExitCode int
	Output        io.Writer
	Error         io.Writer
	indention     uint
}

// Step is a function that returns a sub-task for of the receiving Task.
func (t Task) Step() *Task {
	return &Task{
		ErrorExitCode: t.ErrorExitCode,
		Output:        t.Output,
		Error:         t.Error,
		indention:     t.indention + 1,
	}
}

// Track is a function that logs the tracked status of a given workload function.
func (t Task) Track(description string, workload func() error) error {
	_, _ = fmt.Fprintf(t.Output, t.logTemplate("%s...", false), description)
	if err := workload(); err != nil {
		_, _ = fmt.Fprintln(t.Output, styleError(" Failed"))
		return err
	}

	_, _ = fmt.Fprintln(t.Output, styleDone(" Done"))
	return nil
}
