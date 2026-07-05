package project

import (
	"errors"
	"os/exec"

	"github.com/spenserblack/go-gitutil"
)

// rootFinder is a type that tries to find the root of a project.
type rootFinder interface {
	// Root gets the root of the project.
	//
	// When it would have called an executable, and the executable's path is unsafe, it can return
	// [exec.ErrDot].
	Root() (string, error)
}

// Root gets the root of a project. If the [git.Git] is nil, it will not use git to help find the
// root.
func Root(git gitutil.Git) (string, error) {
	var g rootFinder
	if git != nil {
		g = gitFinder{git: git}
	}
	return root(g, cwd{})
}

// root is the implementation of [Root].
func root(git rootFinder, cwd rootFinder) (string, error) {
	if git != nil {
		root, err := git.Root()
		if errors.Is(err, exec.ErrDot) {
			return "", err
		} else if err == nil {
			return root, nil
		}
	}
	return cwd.Root()
}
