package cmd

import blinkycmd "github.com/Hayao0819/ayaka/cmd/blinky"

func init() {
	subCmds = append(subCmds, blinkycmd.Root())
}
