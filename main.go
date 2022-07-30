package main

import (
	"os"
	"troy/src"
	"troy/src/dasm/enum"
	"troy/src/eth"
	"troy/src/ui"
)

func main() {
	args := src.ParseArgs(os.Args)
	config := src.InitConfig(args)
	contract := eth.InitInitialContract(args, config)

	enumerator := enum.NewContractEnum(contract)
	enumerator.Enumerate()

	ui.Start(enumerator, config)
}
