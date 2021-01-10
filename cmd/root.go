package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kata",
	Short: "kata",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("Subcommand required")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
