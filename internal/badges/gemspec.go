package badges

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spenserblack/autobadges/internal/stringutil"
)

// GemSpec will generate a badge.fury.io badge for each Ruby Gem. Will return an empty or nil slice
// if there aren't any gemspec files in the expected format.
func GemSpec(root string) []string {
	pattern := filepath.Join(root, "*.gemspec")
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil
	}
	badges := make([]string, 0, len(matches))

	for _, gemspec := range matches {
		raw, err := os.ReadFile(gemspec)
		if err != nil {
			continue
		}
		name := gemSpecNameBytes(raw)
		if name == "" {
			continue
		}

		badge := fmt.Sprintf("[![Gem Version](https://badge.fury.io/rb/%[1]s.svg)](https://badge.fury.io/rb/%[1]s)", name)
		badges = append(badges, badge)
	}

	return badges
}

// gemSpecNameBytes reads bytes representing a .gemspec file and attempts to get the spec's name.
func gemSpecNameBytes(b []byte) string {
	v := gemSpecVarBytes(b)
	if v == "" {
		return ""
	}
	rawPattern := fmt.Sprintf(`\s+%s\.name\s*=\s*(.+)\n`, regexp.QuoteMeta(v))
	pattern, err := regexp.Compile(rawPattern)
	if err != nil {
		return ""
	}
	matches := pattern.FindSubmatch(b)
	if len(matches) == 0 {
		return ""
	}

	quoted := string(matches[1])
	unquoted := stringutil.Unquote(quoted)
	if unquoted == quoted {
		// NOTE Not a valid string
		return ""
	}
	return unquoted
}

var gemSpecVariablePattern = regexp.MustCompile(`(?:^|\s)Gem::Specification\.new\s+do\s+\|([A-Za-z]\w*)\|`)

// gemSpecVarBytes reads bytes representing a .gemspec file and attempts to get the variable name that represents the spec.
func gemSpecVarBytes(b []byte) string {
	matches := gemSpecVariablePattern.FindSubmatch(b)
	if len(matches) != 0 {
		return string(matches[1])
	}

	return ""
}
