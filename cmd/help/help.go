package help

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Help = &cobra.Command{
		Use:   "help",
		Short: "Get info of commands",
		Long:  "Get information list and informaition about commands in this",
		Run:   getHelp,
	}
)

func getHelp(cmd *cobra.Command, arg []string) {

	fmt.Println("Help")
}
