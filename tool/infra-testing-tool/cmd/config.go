package cmd

import (
	"github.com/spf13/cobra"
)

var (
	configCmd = &cobra.Command{
		Use: "config",
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}
)
