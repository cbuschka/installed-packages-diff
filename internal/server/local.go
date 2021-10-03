package server

import (
	"os"
	"os/exec"
)

type LocalChannel struct{}

func (local *LocalChannel) IsFile(path string) (bool, error) {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return stat.Mode().IsRegular(), nil
}

func (local *LocalChannel) Run(command ...string) (string, string, error) {

	out, err := exec.Command(command[0], command[1:]...).Output()
	if exiterr, ok := err.(*exec.ExitError); ok {
		return "", "", &ExitError{ExitCode: exiterr.ExitCode()}
	}

	if err != nil {
		return "", "", err
	}

	return string(out), "", nil
}
