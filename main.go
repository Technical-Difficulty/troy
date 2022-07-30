package main

import (
	"os"
	"troy/src"
	"troy/src/eth"
	"troy/src/ui"
)

func main() {
	args := src.ParseArgs(os.Args)
	config := src.InitConfig(args)
	contract := eth.InitInitialContract(args, config)

	ui.Start(contract, config)
}
