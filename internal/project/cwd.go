package project

import "os"

// cwd gets the root from the working directory.
type cwd struct{}

func (cwd) Root() (string, error) {
	return os.Getwd()
}
