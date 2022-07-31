package enum

import (
	"troy/src/dasm/enum/functions"
	"troy/src/dasm/enum/instructions"
	"troy/src/eth"
)

type ContractAnalysis struct {
	Contract     eth.Contract
	SignatureMap map[uint64][]string // JUMPDEST => func signatures
	JumpMap      map[uint64]uint64   // Jump => Location (PC => PC)
}

func NewContractAnalysis(contract eth.Contract) ContractAnalysis {
	return ContractAnalysis{
		Contract: contract,
	}
}

// todo: implement generic enumerator type with enumerate() func
func (e *ContractAnalysis) Enumerate() {
	functionEnum := functions.NewFunctionEnum(e.Contract.Instructions)
	instructionEnum := instructions.NewInstructionEnum(e.Contract.Instructions)

	e.SignatureMap = functionEnum.EnumerateSignatures()
	e.JumpMap = instructionEnum.EnumerateJumps()
}
