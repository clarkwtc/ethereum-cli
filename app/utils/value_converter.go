package utils

import (
    "github.com/ethereum/go-ethereum/params"
    "math/big"
)

func ToEther(value *big.Int) string {
    return new(big.Float).Quo(new(big.Float).SetInt(value), big.NewFloat(params.Ether)).Text('f', 18)
}

func ToGWei(value *big.Int) string {
    return new(big.Float).Quo(new(big.Float).SetInt(value), big.NewFloat(params.GWei)).Text('f', 9)
}
