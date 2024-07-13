package command

import "fmt"

type HelpCommand struct{}

func (command *HelpCommand) Execute() {
    fmt.Printf("usage: eth [help]\n\n")
    fmt.Printf("Here are common eth commands:\n\n")
    fmt.Printf("%-15s Fetch latest 10 record transactions hex\n", "latest:")
    fmt.Printf("%-15s Find transaction info from transaction hex\n", "transaction:")
    fmt.Printf("%-15s Exit current program\n", "exit:")
}
