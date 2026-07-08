package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spenserblack/autobadges/internal/badges"
	"github.com/spenserblack/autobadges/internal/readme"
	"github.com/spenserblack/go-gitutil"
	"github.com/spf13/cobra"
)

var toTerminal bool

var rootCmd = &cobra.Command{
	Use:   "autobadges [README PATH]",
	Short: "Add badges to your README based on your project files",
	Long: heredoc.Doc(`
		This adds badges to your README file.

		The badges to add are determined by the files in your project. The following are supported:

		- GitHub
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
			path string
			f    *os.File
			git  gitutil.Git
			err  error
		)
		stdout := cmd.OutOrStdout()
		stderr := cmd.ErrOrStderr()

		if len(args) > 0 {
			path = args[0]
			f, err = readme.Open(path)
		} else {
			var cwd string
			cwd, err = os.Getwd()
			if err != nil {
				fmt.Fprintln(stderr, "Couldn't get cwd:", err)
				os.Exit(1)
			}
			git, err = gitutil.New(cwd)
			if err != nil {
				git = nil
			}
			path, f, err = readme.FindAndOpen(git)
		}
		if err != nil {
			return fmt.Errorf("Couldn't open README file: %w", err)
		}

		root := filepath.Dir(path)
		badges := badges.Badges(root, git)

		if toTerminal {
			for _, badge := range badges {
				fmt.Fprintln(stdout, badge)
			}
			return nil
		}

		err = readme.AddBadges(f, badges)

		if err != nil {
			fmt.Fprintln(stderr, "Couldn't add badges to", path)
			fmt.Fprintln(stderr, err)
			os.Exit(1)
		}

		fmt.Fprintln(stdout, "Badges written to", path)

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&toTerminal, "to-terminal", "t", false, "Write the badges to the terminal rather than to your README")
}
