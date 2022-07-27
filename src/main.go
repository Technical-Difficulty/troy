package src

import (
	"encoding/hex"
	"fmt"
	"troy/src/dasm"
	"troy/src/eth"

	log "github.com/sirupsen/logrus"
)

var notableOps = map[byte]bool{
	0xf0: true, // CREATE
	0xf1: true, // CALL
	0xf4: true, // DELEGATECALL
	0xf5: true, // CREATE2
	0xff: true, // SELFDESTRUCT
	0x63: true, // PUSH4
}

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
		if _, ok := notableOps[byte(it.Op())]; ok {
			// TODO: make a seperate function to parse op codes to highlight anything
			// notable and then pipe it into some nice terminal UI. Should be able to
			// do more advanced lookups too like if the opcode is CALL look for future
			// SSTORES. We can also look at the Arg of CALL to highlight interesting
			// function sigs like approve, mint, transfer etc.
			log.Info(fmt.Sprintf("Noteable opcode found: %v", it.Op()))
		}

		if it.Arg() != nil && 0 < len(it.Arg()) {
			fmt.Printf("%05x: %v %#x\n", it.PC(), it.Op(), it.Arg())
		} else {
			fmt.Printf("%05x: %v\n", it.PC(), it.Op())
		}
	}

}
