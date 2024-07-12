package command

import (
    "fmt"
    "geth_cli/app/client"
    "geth_cli/app/events"
    "geth_cli/app/utils"
    "github.com/ethereum/go-ethereum/common"
    "math"
    "math/big"
    "time"
)

type ShowTransactionsCommand struct {
    EthClient *client.EthClient
}

func (command *ShowTransactionsCommand) Execute() {
    transactionHex, _ := utils.NewCommandLine().Input()
    
    if transactionHex == "" {
        return
    }

    ethClient := command.EthClient
    receipts := ethClient.GetTransactionReceipt(transactionHex)
    transaction := ethClient.GetTransactionByHex(transactionHex)

    sender, err := client.GetAddresBySender(transaction)
    if err != nil {
        return
    }

    block, err := ethClient.GetBlockByHash(receipts.BlockHash)
    if err != nil {
        return
    }

    event := &events.TransactionEvent{TransactionHash: common.HexToHash(transactionHex), Block: block, Receipts: receipts, Sender: sender, Transaction: transaction}
    showTransactionDetails(event)
}

func showTransactionDetails(event *events.TransactionEvent) {
    receipts := event.Receipts
    tx := event.Transaction
    block := event.Block
    fmt.Println("Transaction Hash:", event.TransactionHash)
    fmt.Println("Status:", receipts.Status != 0)
    fmt.Println("Block:", receipts.BlockNumber)
    utcTime := utils.ToUTCTime(block.Time())
    fmt.Printf("%0.f days ago (%s)\n", math.Floor(time.Since(utcTime).Hours()/24), utcTime.Format("Jan-02-2006 03:04:05 PM UTC"))

    fmt.Println("-------------------------------")
    fmt.Println("From:", event.Sender)
    fmt.Println("To:", tx.To().Hex())

    fmt.Println("-------------------------------")
    fmt.Printf("Value: %s ETH\n", utils.ToEther(tx.Value()))

    fee := new(big.Int).Mul(tx.GasPrice(), big.NewInt(int64(receipts.GasUsed)))
    fmt.Printf("Transaction Fee: %s ETH\n", utils.ToEther(fee))
    fmt.Printf("Gas Price: %s Gwei\n", utils.ToGWei(tx.GasPrice()))

    fmt.Println("-------------------------------")
    fmt.Printf("Txn Type: %d (EIP-1559)\n", tx.Type())
    fmt.Println("Nonce:", tx.Nonce())
    fmt.Println("Position In Block:", receipts.TransactionIndex)
    fmt.Println("Input Data:", tx.Data())
}
