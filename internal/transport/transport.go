package transport

import "fmt"

type ExitError struct {
	ExitCode int
}

func (err *ExitError) Error() string {
	return fmt.Sprintf("exit with code %d", err.ExitCode)
}

type Transport interface {
	ExecCommand(command ...string) (string, string, error)
	IsFile(path string) (bool, error)
}
