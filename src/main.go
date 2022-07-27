package src

import (
	"fmt"
	"troy/src/eth"
)

func Start() {
	config := InitConfig()
	contract := eth.NewContract(config.Address, config.RPCUrl, config.Network)

	fmt.Println(contract.GetByteCode())
}
