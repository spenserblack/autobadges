// Package cmd contains the CLI.
package cmd

import (
	"fmt"
	"os"
)

// Execute executes the CLI.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
