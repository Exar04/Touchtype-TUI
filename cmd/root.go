package cmd

import (
	"github.com/Exar04/touchtype-tui/cmd/highscore"
	"github.com/Exar04/touchtype-tui/cmd/start"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "root",
		Short: "Touch typing",
		Long:  "play and test your typing speed",
	}
)

func Execute() error {
	rootCmd.AddCommand(start.StartCmd)
	rootCmd.AddCommand(highscore.Highscore)
	return rootCmd.Execute()
}
