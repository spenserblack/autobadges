package project

import (
	"os/exec"
	"strings"
)

// git gets the root from the git repository.
type git struct{}

func (git) Root() (string, error) {
	path, err := exec.LookPath("git")
	if err != nil {
		return "", err
	}
	cmd := exec.Command(path, "rev-parse", "--no-revs", "--show-toplevel")
	root, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(root)), nil
}
