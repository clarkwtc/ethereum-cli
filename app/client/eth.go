package client

import (
    "context"
    "fmt"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
    "math/big"
)

type EthClient struct {
    context context.Context
    client  *ethclient.Client
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

func (eth *EthClient) Close() {
    eth.Close()
}

func (eth *EthClient) TransactionReceipt(transactionHex string) (*types.Receipt, error) {
    transactionHash := common.HexToHash(transactionHex)
    receipts, err := eth.client.TransactionReceipt(eth.context, transactionHash)
    if err != nil {
        fmt.Println(err)
    }

    return receipts, err
}

func (eth *EthClient) TransactionByHash(transactionHex string) (*types.Transaction, error) {
    transactionHash := common.HexToHash(transactionHex)
    transaction, pending, err := eth.client.TransactionByHash(eth.context, transactionHash)
    if pending {
        fmt.Println(pending)
    }
    if err != nil {
        fmt.Println(err)
    }

    return transaction, err
}

func (eth *EthClient) GetAddresBySender(tx *types.Transaction) (common.Address, error) {
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

func (eth *EthClient) BlockByHash(hash common.Hash) (*types.Block, error) {
    block, err := eth.client.BlockByHash(eth.context, hash)
    if err != nil {
        fmt.Println(err)
    }

    return block, err
}

func (eth *EthClient) HeaderByNumber() (*types.Header, error) {
    number, err := eth.client.HeaderByNumber(eth.context, nil)
    if err != nil {
        fmt.Println(err)
    }

    return number, err
}

func (eth *EthClient) BlockByNumber(number *big.Int) (*types.Block, error) {
    block, err := eth.client.BlockByNumber(eth.context, number)
    if err != nil {
        fmt.Println(err)
    }

    return block, err
}

func (eth *EthClient) BalanceAt(account common.Address) (*big.Int, error) {
    balance, err := eth.client.BalanceAt(eth.context, account, nil)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }

    return balance, err
}

func (eth *EthClient) NonceAt(account common.Address) (uint64, error) {
    nonce, err := eth.client.NonceAt(eth.context, account, nil)
    if err != nil {
        fmt.Println(err)
        return 0, err
    }

    return nonce, err
}
