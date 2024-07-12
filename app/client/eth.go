package client

import (
    "context"
    "fmt"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
)

type EthClient struct {
    Context context.Context
    Client  *ethclient.Client
    URL     string
}

func NewEthClient(url string) *EthClient {
    ctx := context.Background()
    client, err := ethclient.DialContext(ctx, url)
    if err != nil {
        fmt.Println(err)
    }
    return &EthClient{ctx, client, url}
}

func (eth *EthClient) GetTransactionReceipt(transactionHex string) *types.Receipt {
    transactionHash := common.HexToHash(transactionHex)
    receipts, err := eth.Client.TransactionReceipt(eth.Context, transactionHash)
    if err != nil {
        fmt.Println(err)
    }

    return receipts
}

func (eth *EthClient) GetTransactionByHex(transactionHex string) *types.Transaction {
    transactionHash := common.HexToHash(transactionHex)
    transaction, pending, err := eth.Client.TransactionByHash(eth.Context, transactionHash)
    if pending {
        fmt.Println(pending)
    }
    if err != nil {
        fmt.Println(err)
    }

    return transaction
}

func GetAddresBySender(tx *types.Transaction) (common.Address, error) {
    var signer types.Signer
    switch {
    case tx.Type() == types.AccessListTxType:
        signer = types.NewEIP2930Signer(tx.ChainId())
    case tx.Type() == types.DynamicFeeTxType:
        signer = types.NewLondonSigner(tx.ChainId())
    default:
        signer = types.NewEIP155Signer(tx.ChainId())
    }

    return types.Sender(signer, tx)
}

func (eth *EthClient) GetBlockByHash(hash common.Hash) (*types.Block, error) {
    block, err := eth.Client.BlockByHash(eth.Context, hash)
    if err != nil {
        fmt.Println(err)
    }

    return block, err
}
