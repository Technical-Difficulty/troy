package dasm

import "github.com/ethereum/go-ethereum/core/vm"

type Instruction struct {
	PC      uint64
	OpCode  vm.OpCode
	Operand []byte
}
