package badges

import "testing"

func TestGemSpecFromPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "path is github-linguist.gemspec",
			path: "github-linguist.gemspec",
			want: "[![Gem Version](https://badge.fury.io/rb/github-linguist.svg)](https://badge.fury.io/rb/github-linguist)",
		},
		{
			name: "path is path/to/github-linguist.gemspec",
			path: "path/to/github-linguist.gemspec",
			want: "[![Gem Version](https://badge.fury.io/rb/github-linguist.svg)](https://badge.fury.io/rb/github-linguist)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gemSpecFromPath(tt.path); got != tt.want {
				t.Fatalf(`gemSpecFromPath() = %q, want %q`, got, tt.want)
			}
		})
	}
}
