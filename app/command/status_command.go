package command

import (
    "fmt"
    "geth_cli/app/configure"
)

type StatusCommand struct {
    Config *configure.Config
}

func (command *StatusCommand) Execute() {
    fmt.Printf("Current connecting url: %s \n", command.Config.URL)
}
