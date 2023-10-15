package main

import (
	"fmt"

	"github.com/Exar04/touchtype-tui/cmd"
)

func main() {
	err := cmd.Execute()

	if err != nil {
		fmt.Println(err)
	}
}
