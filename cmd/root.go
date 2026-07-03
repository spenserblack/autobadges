package cmd

import (
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spenserblack/autobadges/internal/readme"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "autobadges [README PATH]",
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
			path string
			f    *os.File
			err  error
		)
		stdout := cmd.OutOrStdout()
		stderr := cmd.ErrOrStderr()

		if len(args) > 0 {
			path = args[0]
			f, err = readme.Open(path)
		} else {
			path, f, err = readme.FindAndOpen()
		}
		if err != nil {
			return fmt.Errorf("Couldn't open README file: %w", err)
		}

		// readLimit is the number of bytes to read. If we can't find a title within a whole
		// kilobyte of text, 🤷
		const readLimit int = 1 << 10
		bytes := make([]byte, readLimit)
		n, err := f.Read(bytes)
		if err != nil {
			fmt.Fprintln(stderr, "Couldn't read", path)
			fmt.Fprintln(stderr, err)
			os.Exit(1)
		}
		bytes = bytes[:n]
		index := readme.TitleEndIndex(bytes)
		fmt.Fprintln(stdout, "Badges should be written at index", index)

		return nil
	},
}
