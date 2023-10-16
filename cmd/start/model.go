package start

import tea "github.com/charmbracelet/bubbletea"

type typingmodel struct {
	TargetWord string
	TypedWord  string
}

type message string

const (
	TypeKey message = "type"
	QuitKey message = "quit"
)

func (m typingmodel) Init() tea.Cmd {
	return nil
}
