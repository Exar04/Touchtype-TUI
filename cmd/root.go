package cmd

import (
	"github.com/Exar04/touchtype-tui/cmd/help"
	"github.com/Exar04/touchtype-tui/cmd/highscore"
	"github.com/Exar04/touchtype-tui/cmd/start"
	"github.com/Exar04/touchtype-tui/cmd/lelo"
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
	rootCmd.AddCommand(help.Help)
	rootCmd.AddCommand(lelo.Lelo)
	return rootCmd.Execute()
}
