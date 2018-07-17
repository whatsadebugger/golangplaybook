package main

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink/store/assets"
)

type WithdrawalRequest struct {
	Address common.Address `json:"address"`
	Amount  assets.Link    `json:"amount"`
}

func main() {
	var wr WithdrawalRequest
	json.Unmarshal([]byte(`{"address": "0xBEEF0002200AA000000000000000000000AA0000", "amount": "1000000"}`), &wr)
	fmt.Printf("%+v\n", wr)
}
