package core

import (
    "geth_cli/app/client"
    "geth_cli/app/command"
    "geth_cli/app/utils"
    "github.com/joho/godotenv"
    "log"
    "os"
)

type CLI struct {
    controller *Controller
}

func NewCLI() *CLI {
    initEnviorment()
    return &CLI{initContorller()}
}

func initEnviorment() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalln("Loading env fail")
    }
}

func initContorller() *Controller {
    url := os.Getenv("BLOCK_CHAIN_URL")
    ethClient := client.NewEthClient(url)

    controller := NewController()
    controller.AddCommand("help", &command.HelpCommand{})
    controller.AddCommand("exit", &command.ExitCommand{})
    controller.AddCommand("transaction", &command.ShowTransactionDetailCommand{EthClient: ethClient})
    controller.AddCommand("latest", &command.ShowLatestTransactionCommand{EthClient: ethClient})
    return controller
}

func (cli *CLI) Run() {
    commandLine := utils.NewCommandLine()
    for {
        input, _ := commandLine.Input()
        cli.controller.Execute(input)
    }
}
