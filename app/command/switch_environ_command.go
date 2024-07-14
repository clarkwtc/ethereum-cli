package command

import (
    "fmt"
    "geth_cli/app/client"
    "geth_cli/app/configure"
    "geth_cli/app/utils"
)

type SwitchEnvironCommand struct {
    EthClientManager *client.EthClientManager
    Config           *configure.Config
}

func (command *SwitchEnvironCommand) Execute() {
    fmt.Print("Please input url type: ")
    urlType, _ := utils.NewCommandLine().Input()
    url := configure.GetURLByType(urlType)

    command.EthClientManager.ReconnectEthClient(url)
    command.Config.URL = url

    fmt.Printf("Connected ... %s\n", url)
}
