package project

import (
	"strings"

	"github.com/spenserblack/go-gitutil"
)

// gitFinder gets the root from the gitFinder repository.
type gitFinder struct {
	git gitutil.Git
}

func (g gitFinder) Root() (string, error) {
	raw, err := g.git.Output("rev-parse", "--no-revs", "--show-toplevel")
	if err != nil {
		return "", nil
	}
	return strings.TrimSpace(string(raw)), nil
}
