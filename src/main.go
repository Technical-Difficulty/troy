package src

import (
	"encoding/hex"
	"troy/src/dasm"
	"troy/src/eth"
	"troy/src/ui"

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

	instructions := dasm.GetInstructions(code)

	// Returns function signature as enum.FuncSig
	// fs.String() 0xfdf80bda
	// fs.Lookup() "transfer(address,uint256)"
	// enum.FuncSigs(instructions)

	ui.Start(instructions)
}
