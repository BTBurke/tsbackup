package main

import (
	"fmt"
	"github.com/BTBurke/tsbackup/msgs"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(msgs.NoArgsHelp)
	}

	if len(os.Args) > 1 {
		cmd := os.Args[1]
		rest := os.Args[1 : len(os.Args)-1]
		switch strings.ToLower(cmd) {
		case "help":
			fmt.Println(msgs.Help)
		case "add":
			commands.Add(rest)
		default:
			fmt.Println(msgs.UnknownCmd(os.Args[1]))
		}
	}
}
