package command

import (
    "fmt"
    "geth_cli/app/client"
    "geth_cli/app/utils"
    "github.com/ethereum/go-ethereum/accounts/keystore"
    "os"
)

type ShowAccountCommand struct {
    EthClient *client.EthClient
}

func (command *ShowAccountCommand) Execute() {
    fmt.Println("Please input keystore path:")
    fileName, err := utils.NewCommandLine().Input()
    if err != nil {
        fmt.Println(err)
        return
    }
    
    keyJson, err := os.ReadFile("./keystore/" + fileName)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    fmt.Println("Please input password:")
    password, err := utils.NewCommandLine().Input()
    if err != nil {
        fmt.Println(err)
        return
    }
    
    privateKey, err := keystore.DecryptKey(keyJson, password)
    if err != nil {
        fmt.Println(err)
        return
    }

    address := privateKey.Address
    ethClient := command.EthClient

    balance, err := ethClient.BalanceAt(address)
    if err != nil {
        return
    }

    nonce, err := ethClient.NonceAt(address)
    if err != nil {
        return
    }

    ether := utils.ToEther(balance)
    fmt.Printf("Address: %s\n", address.Hex())
    fmt.Printf("Balance: %s ETH\n", ether)
    fmt.Printf("Nonce: %d\n", nonce)
}
