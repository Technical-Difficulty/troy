package src

import (
	"encoding/hex"
	"fmt"
	"strings"
	"troy/src/dasm"
	"troy/src/eth"
)

func Start() {
	config := InitConfig()
	contract := eth.NewContract(config.Address, config.RPCUrl, config.Network)
	code := contract.GetByteCode()
	code = strings.Replace(code, "0x", "", -1)

	fmt.Printf("%v\n", code)

	script, err := hex.DecodeString(code)
	if err != nil {
		fmt.Println("Error decoding")
	}

	it := dasm.NewInstructionIterator(script)
	for it.Next() {
		if it.Arg() != nil && 0 < len(it.Arg()) {
			fmt.Printf("%05x: %v %#x\n", it.PC(), it.Op(), it.Arg())
		} else {
			fmt.Printf("%05x: %v\n", it.PC(), it.Op())
		}
	}
}
