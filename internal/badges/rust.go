package badges

import (
	"fmt"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// Cargo will generate a crates.io version badge and a docs.rs documentation badge if a
// Cargo.toml file is found and read. This will be returned as a 2-length string slice, or a nil
// slice.
func Cargo(root string) []string {
	path := filepath.Join(root, "Cargo.toml")
	metadata := struct {
		Package struct{
			Name string
		}
	}{}
	_, err := toml.DecodeFile(path, &metadata)
	name := metadata.Package.Name
	if name == "" || err != nil {
		return nil
	}
	crates := fmt.Sprintf("[![crates.io](https://img.shields.io/crates/v/%[1]s.svg)](https://crates.io/crates/%[1]s)", name)
	docs := fmt.Sprintf("[![docs.rs](https://docs.rs/%[1]s/badge.svg)](https://docs.rs/%[1]s)", name)
	return []string{crates, docs}
}
