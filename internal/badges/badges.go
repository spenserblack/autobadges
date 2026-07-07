// Package badges provides utilities for generating Markdown badge text.
package badges

import "github.com/spenserblack/go-gitutil"

// Badges gets all of the badges for a project.
func Badges(root string, git gitutil.Git) []string {
	// NOTE 2 Cargo badges
	badges := []string{}
	cargo := Cargo(root)
	badges = append(badges, cargo...)
	badges = append(badges, GoMod(root))
	if remote, err := getRemote(git); err == nil {
		githubWorkflows := GitHubWorkflows(root, remote)
		badges = append(badges, githubWorkflows...)
	}

	return badges
}

// getRemote tries to get the primary remote for the Git repository.
func getRemote(git gitutil.Git) (remote gitutil.Remote, err error) {
	remotes, err := gitutil.RemoteNames(git)
	if err != nil {
		return
	}
	main, err := gitutil.MainRemote(remotes)
	if err != nil {
		return
	}
	url, err := gitutil.RawRemoteUrl(git, main)
	if err != nil {
		return
	}
	return gitutil.ParseRemoteUrl(url)
}
