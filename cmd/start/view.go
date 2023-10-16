package start

import "fmt"

func (m typingmodel) View() string {

	s := fmt.Sprintf("%s\n%s", m.TargetWord, m.TypedWord)
	return s
}
