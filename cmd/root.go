package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "egsql",
}

func exitError(msg interface{}) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

// Execute start command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		exitError(err)
	}
}
