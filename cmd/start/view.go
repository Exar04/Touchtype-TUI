package start

import "fmt"

func (m typingmodel) View() string {
	// Render the TUI interface based on the player's state
	// Return the rendered string
	return fmt.Sprintf("Type the word: %s\nYour input: %s", m.TargetWord, m.TypedWord)
}
