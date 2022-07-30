package enum

import (
	"troy/src/dasm/enum/functions"
	"troy/src/eth"
)

type ContractEnum struct {
	Contract       eth.Contract
	FuncSignatures map[uint64][]string // pc => func signatures
}

func NewContractEnum(contract eth.Contract) ContractEnum {
	return ContractEnum{
		Contract: contract,
	}
}

func (e *ContractEnum) Enumerate() {
	enum := functions.NewFunctionEnum(e.Contract.Instructions)

	e.FuncSignatures = enum.EnumerateSignatures()
}
