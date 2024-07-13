package command

import (
    "fmt"
    "geth_cli/app/client"
    "geth_cli/app/configure"
    "geth_cli/app/utils"
)

type SwitchEnvironCommand struct {
    EthClient *client.EthClient
}

func (command *SwitchEnvironCommand) Execute() {
    fmt.Print("Please input url type: ")
    urlType, _ := utils.NewCommandLine().Input()
    url := configure.GetURLByType(urlType)

    command.EthClient.Client.Close()
    command.EthClient = client.NewEthClient(url)

    fmt.Printf("Connected ... %s\n", url)
}
