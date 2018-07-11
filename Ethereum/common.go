package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"strconv"
)

// this is me trying to figure out how im going to deal with withdrawals

func main() {
	// 1 naz
	naz := "0.000000000000000001"
	v, _ := strconv.ParseFloat(naz, 64)

	fmt.Println(naz, v)
	v = v * 1e18
	fmt.Println(v)
	fmt.Println(common.HexToAddress("0x0").String())
	fmt.Println(common.HexToAddress("0x1").String())
	fmt.Println("common.IsHexAddress 0x0", common.IsHexAddress("0x0"))
	fmt.Println(common.HexToAddress("0x0") == common.HexToAddress("0x0"))
	fmt.Println(common.HexToAddress("0x0") == common.HexToAddress("0x1"))
	fmt.Println(common.HexToAddress("0x0") != common.HexToAddress("0x1"))
}
