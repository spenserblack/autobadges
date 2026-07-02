package project

import (
	"errors"
	"os/exec"
	"testing"
)

func TestRoot(t *testing.T) {
	tests := []struct {
		name string
		git  rootFinder
		cwd  rootFinder
		want string
		err  error
	}{
		{
			name: "Git succeeds to find a root path",
			git:  hardcodedRoot("/tmp/my-project"),
			want: "/tmp/my-project",
		},
		{
			name: "Git fails because it resolved to an executable in the current directory",
			git:  errRoot{exec.ErrDot},
			err:  exec.ErrDot,
		},
		{
			name: "Git fails, but cwd succeeds",
			git:  errRoot{errors.New(":(")},
			cwd:  hardcodedRoot("/home/me/my-project"),
			want: "/home/me/my-project",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := root(tt.git, tt.cwd)
			if !errors.Is(err, tt.err) {
				t.Fatalf(`err = %v, want %v`, err, tt.err)
			}
			if got != tt.want {
				t.Fatalf(`got = %q, want %q`, got, tt.want)
			}
		})
	}
}

// errRoot always fails with the given error.
type errRoot struct {
	// err is the error to fail with.
	err error
}

func (r errRoot) Root() (string, error) {
	return "", r.err
}

// hardcodedRoot always succeeds with the given string.
type hardcodedRoot string

func (r hardcodedRoot) Root() (string, error) {
	return string(r), nil
}
