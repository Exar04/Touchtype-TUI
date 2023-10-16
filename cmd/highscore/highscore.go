package highscore

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Highscore = &cobra.Command{
		Use:   "highscore",
		Short: "Starts typing game",
		Long:  "Your game will be starting shortly",
		Run:   getHighscore,
	}
)

func getHighscore(cmd *cobra.Command, arg []string) {

	fmt.Println("Your high score is: ")
}
