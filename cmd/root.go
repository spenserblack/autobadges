package cmd

import (
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spenserblack/autobadges/internal/project"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "autobadges [PROJECT ROOT PATH]",
	Short: "Add badges to your README based on your project files",
	Long: heredoc.Doc(`
		This adds badges to your README file.

		The badges to add are determined by the files in your project. The following are supported:

		- GitHub
		  - Release download count for GitHub CLI extensions
		  - workflow status
		- Go
		  - Go reference for Go modules that _are not_ CLI extensions
		- JavaScript
		  - NPM release
		- Ruby
		  - Rubygems badge.fury.io
		- Rust
		  - cargo.toml release
		  - docs.rs badge
	`),
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			root string
			err  error
		)
		stdout := cmd.OutOrStdout()

		if len(args) > 0 {
			root = args[0]
		} else {
			root, err = project.Root()
			if err != nil {
				return fmt.Errorf("Couldn't find project root: %w", err)
			}
		}

		fmt.Fprintf(stdout, "root: %q\n", root)

		return nil
	},
}
