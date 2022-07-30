package dasm

import (
	"github.com/ethereum/go-ethereum/core/vm"
)

type (
	Instruction struct {
		PC      uint64
		OpCode  vm.OpCode
		Operand []byte
	}

	InstructionSet struct {
		instructions []Instruction
	}
)

func NewInstructionSet(instructions []Instruction) InstructionSet {
	return InstructionSet{
		instructions: instructions,
	}
}

func (s *InstructionSet) Array() []Instruction {
	return s.instructions
}
