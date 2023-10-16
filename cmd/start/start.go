package start

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var (
	StartCmd = &cobra.Command{
		Use:   "start",
		Short: "Starts typing game",
		Long:  "Your game will be starting shortly",
		Run:   StartTypingGame,
	}
)

func StartTypingGame(cmd *cobra.Command, arg []string) {
	initialModel := typingmodel{
		TargetWord: "programming", // Change this to the word you want users to type
	}
	p := tea.NewProgram(initialModel)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	// reader := bufio.NewReader(os.Stdin)

	// fmt.Println(textgenerator.TextGenerator())

	// input, _ := reader.ReadString('\n')

	// fmt.Print(input)
}
