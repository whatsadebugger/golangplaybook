package main

import (
	"os/exec"
)

func main() {
	xdgOpen("magnet:?xt=urn:btih:5JDFCUZXLLHYNIM3WSOGOSYXXZAZXFSD")
}

// opens a magnet link in deluge-gtk
func xdgOpen(mag string) {
	cmd1 := exec.Command("xdg-open", mag)
	if err := cmd1.Run(); err != nil {
		panic(err)
	}
}
