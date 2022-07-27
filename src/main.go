package src

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"troy/src/dasm"
	"troy/src/eth"
)

func Start() {
	config := InitConfig()
	contract := eth.NewContract(config.Address, config.RPCUrl, config.Network)
	code := contract.GetByteCode()

	log.WithField("Byte Code", code).Info("Retrieved byte code successfully")

	it := dasm.NewInstructionIterator(code)
	for it.Next() {
		if it.Arg() != nil && 0 < len(it.Arg()) {
			fmt.Printf("%05x: %v %#x\n", it.PC(), it.Op(), it.Arg())
		} else {
			fmt.Printf("%05x: %v\n", it.PC(), it.Op())
		}
	}
}
