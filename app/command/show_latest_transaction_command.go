package command

import (
    "fmt"
    "geth_cli/app/client"
)

type ShowLatestTransactionCommand struct {
    EthClientManager *client.EthClientManager
}

func (command *ShowLatestTransactionCommand) Execute() {
    ethClientManager := command.EthClientManager

    header, err := ethClientManager.HeaderByNumber()
    if err != nil {
        return
    }

    block, err := ethClientManager.BlockByNumber(header.Number)
    if err != nil {
        return
    }

    for index, tx := range block.Transactions()[:10] {
        fmt.Printf("Transaction hex %d: %s\n", index+1, tx.Hash())
    }
}
