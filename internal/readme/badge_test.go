package readme

import (
	"io"
	"os"
	"testing"

	"github.com/MakeNowJust/heredoc/v2"
)

func TestAddBadges(t *testing.T) {
	tests := []struct {
		name    string
		content string
		badges  []string
		want    string
	}{
		{
			name: "It adds badges under the header",
			content: heredoc.Doc(`
				# My project header

				This is my project description!
			`),
			badges: []string{"![badge](https://img.shields.io/badge/badge-exists-blue)"},
			want: heredoc.Doc(`
				# My project header

				![badge](https://img.shields.io/badge/badge-exists-blue)

				This is my project description!
			`),
		},
		{
			name: "It adds badges to the beginning",
			content: heredoc.Doc(`
				There is no header.

				This is a description.
			`),
			badges: []string{"![badge](https://img.shields.io/badge/badge-exists-blue)"},
			want: heredoc.Doc(`
				![badge](https://img.shields.io/badge/badge-exists-blue)

				There is no header.

				This is a description.
			`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// NOTE Setup
			f, err := os.CreateTemp("", "autobadges-add-badges-test")
			if err != nil {
				panic(err)
			}
			defer os.Remove(f.Name())
			defer f.Close()
			if _, err := f.Write([]byte(tt.content)); err != nil {
				panic(err)
			}
			if _, err := f.Seek(0, io.SeekStart); err != nil {
				panic(err)
			}

			// NOTE Actual test
			err = AddBadges(f, tt.badges)
			if err != nil {
				t.Fatalf(`err = %v, want nil`, err)
			}
			if _, err := f.Seek(0, io.SeekStart); err != nil {
				panic(err)
			}
			gotbytes, err := io.ReadAll(f)
			if err != nil {
				panic(err)
			}
			got := string(gotbytes)
			if got != tt.want {
				t.Errorf(`got = %q, want %q`, got, tt.want)
			}
		})
	}
}
