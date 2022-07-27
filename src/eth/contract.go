package eth

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

func (c *Contract) GetByteCode() string {
	return c.rpc.GetCode(c.Address)
}
