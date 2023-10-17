package start

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m typingmodel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    styledText := ""

    switch msg := msg.(type) {

    case tea.KeyMsg:
        switch msg.String() {
        case "esc":
            return m, tea.Quit
        case "enter": // Enter key
            if strings.TrimSpace(m.TypedWord) == m.TargetWord {
                // m.TypedWord = "" // Clear the typed word
                return m, tea.Quit
            }
        case "backspace":
            m.TypedWord = removeLastCharacter(m.TypedWord)
            m.DisplayedText =  updateTheTypingText(m,styledText)

        default:
            m.TypedWord += msg.String()
            m.DisplayedText =  updateTheTypingText(m,styledText)

        }
    }
    return m, nil
}

func updateTheTypingText(m typingmodel, styledText string) string{
            runes1 := []rune(m.TargetWord)
            runes2 := []rune(m.TypedWord)

            redBoldText := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Bold(true)
            blueText := lipgloss.NewStyle().Foreground(lipgloss.Color("#0000FF"))
            whiteText := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(true)

            for i := 0; i < len(runes1) && i < len(runes2); i++ {
                if runes1[i] == runes2[i] {
                    styledText += blueText.Render(string(runes2[i]))
                } else if runes1[i] != runes2[i] {
                    styledText += redBoldText.Render(string(runes2[i]))
                }
            }
            for i := len(runes2); i < len(runes1); i++ {
               styledText += whiteText.Render(string(runes1[i]))
            }

            return styledText
}

func removeLastCharacter(input string) string {
	if len(input) == 0 {
		return input
	}
	return input[:len(input)-1]
}
