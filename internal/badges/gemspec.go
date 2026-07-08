package badges

import (
	"fmt"
	"path/filepath"
	"strings"
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

	for _, path := range matches {
		badge := gemSpecFromPath(path)
		badges = append(badges, badge)
	}

	return badges
}

// gemSpecFromPath creates a badge.fury.io based on the path to a .gemspec file.
func gemSpecFromPath(path string) string {
	base := filepath.Base(path)
	name := strings.TrimSuffix(base, ".gemspec")
	return fmt.Sprintf("[![Gem Version](https://badge.fury.io/rb/%[1]s.svg)](https://badge.fury.io/rb/%[1]s)", name)
}
