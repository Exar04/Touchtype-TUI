package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "Type",
		Short: "Touch typing",
		Long:  "play and test your typing speed",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
