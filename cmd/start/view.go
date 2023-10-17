package start

import (

	"github.com/charmbracelet/lipgloss"
)

func (m typingmodel) View() string{

    para := lipgloss.NewStyle().BorderForeground(lipgloss.Color("12")).BorderStyle(lipgloss.DoubleBorder()).Padding(3).Width(60)
    para.Align(lipgloss.Center)


//	s := fmt.Sprintf("%s\n%s", m.TargetWord, m.TypedWord)
	return( 
    lipgloss.Place(
		20,
		20,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Left,
			para.Render(m.DisplayedText),
		),
	))
}
