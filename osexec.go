package main

import (
	"github.com/smartcontractkit/chainlink/logger"
	"os/exec"
)

func main() {
	openDeluge("magnet:?xt=urn:btih:5JDFCUZXLLHYNIM3WSOGOSYXXZAZXFSD")
}

// opens a magnet link in deluge-gtk
func openDeluge(mag string) {
	cmd1 := exec.Command("deluge-gtk", mag)
	if err := cmd1.Run(); err != nil {
		logger.Error(err)
	}
}
