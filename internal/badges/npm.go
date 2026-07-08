package badges

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
)

// Npm will generate a badge for the NPM package from a package.json file. Will return an empty
// string if the information is not available.
func Npm(root string) string {
	path := filepath.Join(root, "package.json")
	f, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer f.Close()
	metadata := struct {
		Name string
	}{}
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&metadata)
	if err != nil {
		return ""
	}
	escaped := url.PathEscape(metadata.Name)
	return fmt.Sprintf("[![NPM Version](https://img.shields.io/npm/v/%s)](https://www.npmjs.com/package/%s)", escaped, metadata.Name)
}
