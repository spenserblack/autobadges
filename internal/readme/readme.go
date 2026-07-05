// Package readme provides utilities for interacting with README files.
package readme

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/spenserblack/autobadges/internal/project"
	"github.com/spenserblack/go-gitutil"
)

// Open opens a file in read and write mode.
func Open(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_RDWR, 0)
}

// ErrCannotOpenReadme is raised when an openable README could not be found.
var ErrCannotOpenReadme = errors.New("Could not find a README that could be opened")

// commonReadmeNames is a collection of common README names to try.
var commonReadmeNames = []string{
	"README.md",
	"Readme.md",
	"readme.md",
}

// FindAndOpen tries to find and open the README file for the project. The file is opened in read and
// write mode. The found path and the opened file are returned.
func FindAndOpen(git gitutil.Git) (string, *os.File, error) {
	root, err := project.Root(git)
	if err != nil {
		return "", nil, err
	}

	for _, name := range commonReadmeNames {
		path := filepath.Join(root, name)
		f, err := Open(path)
		if err == nil {
			return path, f, nil
		}
	}

	return "", nil, ErrCannotOpenReadme
}
