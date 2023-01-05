package main

import (
	"fmt"
	"os"

	"github.com/cuotos/gotracks/cmd"
	"github.com/spf13/cobra"
)

var (
	version = "unset"
	commit  = "unset"
)

var rootCmd = cobra.Command{
	Use:     "gotracks",
	Version: fmt.Sprintf("%s-%s", version, commit),
}

func main() {
	rootCmd.AddCommand(cmd.TabCmd)

	rootCmd.CompletionOptions = cobra.CompletionOptions{
		DisableDefaultCmd: true,
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
