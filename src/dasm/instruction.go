package dasm

import (
	"errors"
	"github.com/ethereum/go-ethereum/core/vm"
	log "github.com/sirupsen/logrus"
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

	InstructionMap struct {
		indexToInstruction map[int]Instruction
		pcToIndex          map[uint64]int
	}

	MappedEntry struct {
		Index       int
		Instruction Instruction
	}
)

func NewInstructionSet(instructions []Instruction) InstructionSet {
	return InstructionSet{
		instructions: instructions,
	}
}

func NewInstructionMap() InstructionMap {
	return InstructionMap{
		indexToInstruction: make(map[int]Instruction),
		pcToIndex:          make(map[uint64]int),
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

func (m *InstructionMap) Add(index int, ins Instruction) {
	m.indexToInstruction[index] = ins
	m.pcToIndex[ins.PC] = index
}

func (m *InstructionMap) GetInstruction(index int) (ins Instruction) {
	if ins, found := m.indexToInstruction[index]; found {
		return ins
	}

	log.Fatal("no instruction exists at that index in the instruction map")
	return
}

func (m *InstructionMap) GetIndex(pc uint64) (index int) {
	if index, found := m.pcToIndex[pc]; found {
		return index
	}

	log.Fatal("no instruction with that PC exists in the instruction map")
	return
}

func (m *InstructionMap) GetFromPC(pc uint64) (entry MappedEntry) {
	index := m.GetIndex(pc)
	ins := m.GetInstruction(index)

	return MappedEntry{
		Index:       index,
		Instruction: ins,
	}
}
