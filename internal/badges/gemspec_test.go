package badges

import (
	"testing"

	"github.com/MakeNowJust/heredoc/v2"
)

func TestGemSpecNameBytes(t *testing.T) {
	tests := []struct{
		name string
		content string
		want string
	}{
		{
			name: "spec.name is set",
			content: heredoc.Doc(`
				# frozen_string_literal: true

				Gem::Specification.new do |spec|
				spec.name                   = 'foo'
				spec.version                = '0.1.0'
			`),
			want: "foo",
		},
		{
			name: "double quotes",
			content: heredoc.Doc(`
				# frozen_string_literal: true

				Gem::Specification.new do |spec|
				spec.name                   = "foo-bar"
				spec.version                = '0.1.0'
			`),
			want: "foo-bar",
		},
		{
			name: "double quotes",
			content: heredoc.Doc(`
				# frozen_string_literal: true

				Gem::Specification.new do |spec|
				spec.name                   = "foo-bar"
				spec.version                = '0.1.0'
			`),
			want: "foo-bar",
		},
		{
			name: "spec variable is s",
			content: heredoc.Doc(`
				# frozen_string_literal: true

				Gem::Specification.new do |s|
				s.name                   = "foo-bar"
				s.version                = '0.1.0'
			`),
			want: "foo-bar",
		},
		{
			name: "name is not set",
			content: heredoc.Doc(`
				# frozen_string_literal: true

				Gem::Specification.new do |s|
				s.version                = '0.1.0'
			`),
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gemSpecNameBytes([]byte(tt.content)); got != tt.want {
				t.Fatalf(`gemSpecNameBytes() = %q, want %q`, got, tt.want)
			}
		})
	}
}

func TestGemSpecVarBytes(t *testing.T) {
	tests := []struct{
		name string
		content string
		want string
	}{
		{
			name: "variable is called spec",
			content: heredoc.Doc(`
				# frozen_string_literal: true

				Gem::Specification.new do |spec|
				spec.name                   = 'foo'
				spec.version                = '0.1.0'
			`),
			want: "foo",
		},
		{
			name: "spec variable is s",
			content: heredoc.Doc(`
				# frozen_string_literal: true

				Gem::Specification.new do |s|
				s.name                   = "foo-bar"
				s.version                = '0.1.0'
			`),
			want: "foo-bar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gemSpecNameBytes([]byte(tt.content)); got != tt.want {
				t.Fatalf(`gemSpecNameBytes() = %q, want %q`, got, tt.want)
			}
		})
	}
}
