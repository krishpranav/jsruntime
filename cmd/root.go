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

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version.",
		Run: func(cmd *cobra.Command, args []string) {
			lib.PrintVersion()
		},
	}

	var fsFlag bool
	var netFlag bool
	var envFlag bool

	var runCmd = &cobra.Command{
		Use:   "run [file]",
		Short: "Run a JavaScript file.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			lib.RunFile(args[0])
		},
	}
	runCmd.Flags().BoolVar(&fsFlag, "fs", false, "Allow file system access")
	runCmd.Flags().BoolVar(&netFlag, "net", false, "Allow net access")
	runCmd.Flags().BoolVar(&envFlag, "env", false, "Allow Environment Variables access")

	rootCmd.AddCommand(runCmd, versionCmd)

	return rootCmd
}
