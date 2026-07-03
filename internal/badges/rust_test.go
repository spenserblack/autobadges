package badges

import (
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/MakeNowJust/heredoc/v2"
)

func TestCargo(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    []string
	}{
		{
			name: "package.name is set",
			content: heredoc.Doc(`
				[package]
				name = "autobadges"
			`),
			want: []string{
				"[![crates.io](https://img.shields.io/crates/v/autobadges.svg)](https://crates.io/crates/autobadges)",
				"[![docs.rs](https://docs.rs/autobadges/badge.svg)](https://docs.rs/autobadges)",
			},
		},
		{
			name: "package medatadata is not set",
			content: heredoc.Doc(`
				[foo]
				bar = true
			`),
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// NOTE Setup
			dir := os.TempDir()
			path := filepath.Join(dir, "Cargo.toml")
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
			got := Cargo(dir)
			if !slices.Equal(got, tt.want) {
				t.Fatalf(`Cargo() = %v, want %v`, got, tt.want)
			}
		})
	}
}
