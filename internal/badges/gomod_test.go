package badges

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/MakeNowJust/heredoc/v2"
)

func TestGoMod(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    string
	}{
		{
			name: "module is set",
			content: heredoc.Doc(`
				module github.com/golang/go
			`),
			want: "[![Go Reference](https://pkg.go.dev/badge/github.com/golang/go.svg)](https://pkg.go.dev/github.com/golang/go)",
		},
		{
			name: "module is commented out",
			content: heredoc.Doc(`
				// module github.com/golang/go
			`),
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// NOTE Setup
			dir := os.TempDir()
			path := filepath.Join(dir, "go.mod")
			f, err := os.Create(path)
			if err != nil {
				panic(err)
			}
			defer os.Remove(f.Name())
			defer f.Close()
			if _, err := f.Write([]byte(tt.content)); err != nil {
				panic(err)
			}

			// NOTE Test
			got := GoMod(dir)
			if got != tt.want {
				t.Fatalf(`GoMod() = %q, want %q`, got, tt.want)
			}
		})
	}
}
