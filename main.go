package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/cuotos/gotracks/cmd"
	"github.com/spf13/cobra"
)

var (
	version = "unset"
	commit  = "unset"
)

func getVersionString() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		return info.Main.Version
	}

	return fmt.Sprintf("%s-%s", version, commit)
}

var rootCmd = cobra.Command{
	Use:     "gotracks",
	Version: getVersionString(),
}

func main() {
	rootCmd.AddCommand(cmd.NewTabCmd())

	rootCmd.CompletionOptions = cobra.CompletionOptions{
		DisableDefaultCmd: true,
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
