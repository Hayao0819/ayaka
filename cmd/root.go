package cmd

import "github.com/spf13/cobra"

var subCmds = []*cobra.Command{}

func rootCmd() *cobra.Command {
	cmd := cobra.Command{}

	cmd.AddCommand(subCmds...)

	return &cmd
}

func Execute() error {
	return rootCmd().Execute()
}
