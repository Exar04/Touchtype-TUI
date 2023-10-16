package start

import (
	"regexp"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m typingmodel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	styledText := ""
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "enter": // Enter key
			if strings.TrimSpace(stripANSIColorCodes(m.TypedWord)) == m.TargetWord {
				// m.TypedWord = "" // Clear the typed word
				return m, tea.Quit
			}
		default:

			m.TypedWord += msg.String()
			runes1 := []rune(m.TargetWord)
			runes2 := []rune(stripANSIColorCodes(m.TypedWord))

			redBoldText := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Bold(true)

			blueText := lipgloss.NewStyle().Foreground(lipgloss.Color("#0000FF"))

			for i := 0; i < len(runes1) && i < len(runes2); i++ {
				if runes1[i] == runes2[i] {
					styledText += blueText.Render(string(runes2[i]))
					// m.TypedWord += blueText.Render(string(runes2[i]))
				} else if runes1[i] != runes2[i] {
					styledText += redBoldText.Render(string(runes2[i]))
					// m.TypedWord += redBoldText.Render(string(runes2[i]))
				} else {
					styledText += string(runes2[i])
					// m.TypedWord += string(runes2[i])
				}
			}
			m.TypedWord = styledText

		}
	}
	return m, nil
}

func stripANSIColorCodes(input string) string {
	ansiColorRegex := regexp.MustCompile("\x1b\\[[0-9;]+m")
	return ansiColorRegex.ReplaceAllString(input, "")
}
