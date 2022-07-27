package src

import (
	"encoding/hex"
	"fmt"
	"troy/src/dasm"
	"troy/src/eth"

	log "github.com/sirupsen/logrus"
)

func Start() {
	var code []byte

	config := InitConfig()

	if len(config.Address) > 0 {
		contract := eth.NewContract(config.Address, config.RPCUrl, config.Network)
		code = contract.GetByteCode()
		if len(code) <= 0 {
			log.Error("Failed to retrieved byte code")
			return
		}
	} else {
		script, err := hex.DecodeString(config.Code)
		if err != nil {
			log.Error(err)
			return
		}
		code = script
	}

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
