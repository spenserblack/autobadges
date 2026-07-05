package badges

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var goModuleRe = regexp.MustCompile(`(?:^|\n)module\s+([^\s]+)`)

// GoMod will generate a Go Reference badge if a go.mod file is found and read. Will return an
// empty string if the information is not available.
func GoMod(root string) string {
	path := filepath.Join(root, "go.mod")
	raw, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	captures := goModuleRe.FindSubmatch(raw)
	if len(captures) == 0 {
		return ""
	}

	name := string(captures[1])

	return fmt.Sprintf("[![Go Reference](https://pkg.go.dev/badge/%[1]s.svg)](https://pkg.go.dev/%[1]s)", name)
}
