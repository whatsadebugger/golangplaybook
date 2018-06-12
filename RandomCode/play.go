package main

import (
	"errors"

	"github.com/smartcontractkit/chainlink/logger"
)

func main() {
	err := errors.New("wow bad error")
	logger.Warn("OMG AN ERRRRRROR", err)
}
