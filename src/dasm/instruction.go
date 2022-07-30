package dasm

import (
	"errors"
	"github.com/ethereum/go-ethereum/core/vm"
	"strings"
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

func (s *InstructionSet) ReverseFindInstructionWithPrefix(index int, prefix string) (ins Instruction, err error) {
	if index < 0 {
		return ins, errors.New("index cannot be less than 0")
	}

	if index > len(s.instructions) {
		return ins, errors.New("index cannot be greater than the instructions array length")
	}

	for index > 0 {
		ins = s.instructions[index]

		if strings.HasPrefix(ins.OpCode.String(), prefix) {
			return ins, nil
		}

		index--
	}

	return ins, errors.New("no instruction with that prefix could be found")
}
