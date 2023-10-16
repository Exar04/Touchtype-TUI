package start

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m typingmodel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "enter": // Enter key
			if strings.TrimSpace(m.TypedWord) == m.TargetWord {
				m.TypedWord = "" // Clear the typed word
			}
		default:
			m.TypedWord += msg.String()
		}
	}
	return m, nil
}
