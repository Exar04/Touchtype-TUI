package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	traceCmd = &cobra.Command{
		Use:   "start",
		Short: "Starts typing game",
		Long:  "Your game will be starting shortly",
		Run:   StartTypingGame,
	}
)

func StartTypingGame(cmd *cobra.Command, arg []string) {
    
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type this line")
	input, _ := reader.ReadString('\n')

	fmt.Print(input)
}

func init() {
	rootCmd.AddCommand(traceCmd)
}
