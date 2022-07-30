package enum

import (
	"troy/src/dasm/enum/functions"
	"troy/src/eth"
)

type ContractAnalysis struct {
	Contract     eth.Contract
	SignatureMap map[uint64][]string // pc => func signatures
}

func NewContractAnalysis(contract eth.Contract) ContractAnalysis {
	return ContractAnalysis{
		Contract: contract,
	}
}

func (e *ContractAnalysis) Enumerate() {
	enum := functions.NewFunctionEnum(e.Contract.Instructions)

	e.SignatureMap = enum.EnumerateSignatures()
}
