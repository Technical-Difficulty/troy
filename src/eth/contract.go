package eth

import (
	"encoding/hex"
	log "github.com/sirupsen/logrus"
	"strings"
	"troy/src"
	"troy/src/dasm"
)

type Contract struct {
	Address      string
	ByteCode     []byte
	Instructions []dasm.Instruction
	IsMock       bool // mock contracts are when bytecode is passed as an arg
	rpc          RPC
}

func InitInitialContract(args src.ParsedArgs, config src.Config) Contract {
	if args.Address != "" {
		return NewContract(args.Address, config.RPCUrl, config.Network)
	} else {
		return NewMockContract(args.Code)
	}
}

func NewContract(address string, url string, network string) (contract Contract) {
	contract = Contract{
		Address: address,
		rpc:     NewRPC(url, network),
	}

	bytes := contract.GetByteCode()
	contract.ByteCode = bytes
	contract.Instructions = dasm.GetInstructions(bytes)

	return
}

func NewMockContract(code string) (contract Contract) {
	bytes, err := hex.DecodeString(code)

	if err != nil {
		log.WithField("Byte Code", code).Fatal("Failed to decode provided byte code into bytes.")
	}

	return Contract{
		ByteCode:     bytes,
		Instructions: dasm.GetInstructions(bytes),
		IsMock:       true,
	}
}

func (c *Contract) GetByteCode() []byte {
	if c.IsMock {
		log.Fatal("Cannot retrieve byte code of a mock contract")
	}

	code := c.rpc.GetCode(c.Address)
	code = strings.Replace(code, "0x", "", 1)

	bytes, err := hex.DecodeString(code)
	if err != nil {
		log.WithField("Error", err).
			WithField("Address", c.Address).
			Fatal("Failed to decode retrieved contract byte code")
	}

	return bytes
}
