package command

import (
    "fmt"
    "geth_cli/app/client"
)

type ShowLatestTransactionCommand struct {
    EthClient *client.EthClient
}

func (command *ShowLatestTransactionCommand) Execute() {
    ethClient := command.EthClient

    header, err := ethClient.HeaderByNumber()
    if err != nil {
        return
    }

    block, err := ethClient.BlockByNumber(header.Number)
    if err != nil {
        return
    }

    for index, tx := range block.Transactions()[:10] {
        fmt.Printf("Transaction hex %d: %s\n", index+1, tx.Hash())
    }
}
