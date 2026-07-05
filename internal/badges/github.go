package badges

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spenserblack/go-gitutil"
	"gopkg.in/yaml.v3"
)

// GitHubWorkflows generates a slice of badges for all workflows for the repository.
func GitHubWorkflows(root string, remote gitutil.Remote) []string {
	dir := filepath.Join(root, ".github", "workflows")
	globs := [2]string{"*.yaml", "*.yml"}
	badges := []string{}

	for _, glob := range globs {
		pattern := filepath.Join(dir, glob)
		matches, err := filepath.Glob(pattern)
		if err != nil {
			continue
		}

		for _, workflow := range matches {
			name, display := gitHubWorkflowNames(workflow)
			badge := fmt.Sprintf(
				"[![%[5]s](https://%[1]s/%[2]s/%[3]s/actions/workflows/%[4]s/badge.svg)](https://%[1]s/%[2]s/%[3]s/actions/workflows/%[4]s)",
				remote.Host, remote.Owner, remote.Name, name, display,
			)
			badges = append(badges, badge)
		}
	}

	return badges
}

// gitHubWorkflowName reads a GitHub workflow, getting its name and its display name.
func gitHubWorkflowNames(path string) (name string, display string) {
	name = filepath.Base(path)
	metadata := struct{
		Name string
	}{}
	bytes, err := os.ReadFile(path)
	if err != nil {
		return name, name
	}
	err = yaml.Unmarshal(bytes, &metadata)
	if err != nil {
		return name, name
	}
	display = metadata.Name
	return
}
