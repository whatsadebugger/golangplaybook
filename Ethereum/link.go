package main

import (
	"fmt"
	"github.com/smartcontractkit/chainlink/store/assets"
	"math/big"
)

func main() {
	fmt.Println((*big.Int)(assets.NewLink(1)).Cmp((*big.Int)(assets.NewLink(2))))
	fmt.Println((*big.Int)(assets.NewLink(2)).Cmp((*big.Int)(assets.NewLink(1))))
}
