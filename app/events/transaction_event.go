package events

import (
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
)

type TransactionEvent struct {
    TransactionHash common.Hash
    Block           *types.Block
    Receipts        *types.Receipt
    Sender          common.Address
    Transaction     *types.Transaction
}
