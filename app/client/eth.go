package client

import (
    "context"
    "fmt"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
    "math/big"
)

type EthClientManager struct {
    context context.Context
    client  *ethclient.Client
    URL     string
}

func NewEthClientManager(url string) *EthClientManager {
    ctx := context.Background()
    client, err := ethclient.DialContext(ctx, url)
    if err != nil {
        fmt.Println(err)
    }

    return &EthClientManager{ctx, client, url}
}

func (manager *EthClientManager) ReconnectEthClient(url string) {
    manager.Close()

    ctx := context.Background()
    client, err := ethclient.DialContext(ctx, url)
    if err != nil {
        return
    }

    manager.context = ctx
    manager.client = client
    manager.URL = url
}

func (manager *EthClientManager) Close() {
    manager.client.Close()
}

func (manager *EthClientManager) TransactionReceipt(transactionHex string) (*types.Receipt, error) {
    transactionHash := common.HexToHash(transactionHex)
    receipts, err := manager.client.TransactionReceipt(manager.context, transactionHash)
    if err != nil {
        fmt.Println(err)
    }

    return receipts, err
}

func (manager *EthClientManager) TransactionByHash(transactionHex string) (*types.Transaction, error) {
    transactionHash := common.HexToHash(transactionHex)
    transaction, pending, err := manager.client.TransactionByHash(manager.context, transactionHash)
    if pending {
        fmt.Println(pending)
    }
    if err != nil {
        fmt.Println(err)
    }

    return transaction, err
}

func (manager *EthClientManager) GetAddresBySender(tx *types.Transaction) (common.Address, error) {
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

func (manager *EthClientManager) BlockByHash(hash common.Hash) (*types.Block, error) {
    block, err := manager.client.BlockByHash(manager.context, hash)
    if err != nil {
        fmt.Println(err)
    }

    return block, err
}

func (manager *EthClientManager) HeaderByNumber() (*types.Header, error) {
    number, err := manager.client.HeaderByNumber(manager.context, nil)
    if err != nil {
        fmt.Println(err)
    }

    return number, err
}

func (manager *EthClientManager) BlockByNumber(number *big.Int) (*types.Block, error) {
    block, err := manager.client.BlockByNumber(manager.context, number)
    if err != nil {
        fmt.Println(err)
    }

    return block, err
}

func (manager *EthClientManager) BalanceAt(account common.Address) (*big.Int, error) {
    balance, err := manager.client.BalanceAt(manager.context, account, nil)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }

    return balance, err
}

func (manager *EthClientManager) NonceAt(account common.Address) (uint64, error) {
    nonce, err := manager.client.NonceAt(manager.context, account, nil)
    if err != nil {
        fmt.Println(err)
        return 0, err
    }

    return nonce, err
}
