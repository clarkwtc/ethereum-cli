package command

import (
    "fmt"
    "geth_cli/app/client"
    "geth_cli/app/utils"
    "github.com/ethereum/go-ethereum/accounts/keystore"
    "github.com/ethereum/go-ethereum/crypto"
)

type CreateAccountCommand struct {
    EthClientManager *client.EthClientManager
}

func (command *CreateAccountCommand) Execute() {
    privateKey, err := crypto.GenerateKey()
    if err != nil {
        return
    }

    keyStore := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)

    fmt.Println("Please input password:")
    password, err := utils.NewCommandLine().Input()
    if err != nil {
        return
    }

    account, err := keyStore.ImportECDSA(privateKey, password)
    if err != nil {
        return
    }

    fmt.Printf("Address: %s\n", account.Address.Hex())
    fmt.Printf("Store path: %s\n", account.URL.Path)
}
