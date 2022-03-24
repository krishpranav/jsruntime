package cmd

import (
	"github.com/krishpranav/jsruntime/lib"
	"github.com/spf13/cobra"
)

func Execute() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "Runs the REPL.",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			lib.Repl()
		},
	}

	return rootCmd
}
