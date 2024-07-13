package command

import (
    "fmt"
    "os"
)

type ExitCommand struct{}

func (command *ExitCommand) Execute() {
    fmt.Println("Exit the program ...")
    os.Exit(0)
}
