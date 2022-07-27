package eth

import (
	"encoding/hex"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Contract struct {
	Address string
	rpc     RPC
}

func NewContract(address string, url string, network string) Contract {
	return Contract{
		Address: address,
		rpc:     NewRPC(url, network),
	}
}

func (c *Contract) GetByteCode() []byte {
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
