package lelo

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	Lelo = &cobra.Command{
		Use:   "lelo",
		Short: "Get info of commands",
		Long:  "Get information list and informaition about commands in this",
		Run:   getHelp,
	}
)

func getHelp(cmd *cobra.Command, arg []string) {
    text := "This is test"

    para := lipgloss.NewStyle().BorderForeground(lipgloss.Color("12")).BorderStyle(lipgloss.DoubleBorder()).Padding(3).Width(60)
    para.Align(lipgloss.Center)
    fmt.Println(para.Render(text))

    style := lipgloss.NewStyle().
    Width(100).
    Align(lipgloss.Center).
    Foreground(lipgloss.Color("#FF0000")).Bold(true)

	fmt.Println(style.Render(text))
}
