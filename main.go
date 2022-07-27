package main

import (
	"fmt"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/core/vm"
)

// Iterator for disassembled EVM instructions
type instructionIterator struct {
	code    []byte
	pc      uint64
	arg     []byte
	op      vm.OpCode
	error   error
	started bool
}

// NewInstructionIterator create a new instruction iterator.
func NewInstructionIterator(code []byte) *instructionIterator {
	it := new(instructionIterator)
	it.code = code
	return it
}

// Next returns true if there is a next instruction and moves on.
func (it *instructionIterator) Next() bool {
	if it.error != nil || uint64(len(it.code)) <= it.pc {
		// We previously reached an error or the end.
		return false
	}

	if it.started {
		// Since the iteration has been already started we move to the next instruction.
		if it.arg != nil {
			it.pc += uint64(len(it.arg))
		}
		it.pc++
	} else {
		// We start the iteration from the first instruction.
		it.started = true
	}

	if uint64(len(it.code)) <= it.pc {
		// We reached the end.
		return false
	}

	it.op = vm.OpCode(it.code[it.pc])
	if it.op.IsPush() {
		a := uint64(it.op) - uint64(vm.PUSH1) + 1
		u := it.pc + 1 + a
		if uint64(len(it.code)) <= it.pc || uint64(len(it.code)) < u {
			it.error = fmt.Errorf("incomplete push instruction at %v", it.pc)
			return false
		}
		it.arg = it.code[it.pc+1 : u]
	} else {
		it.arg = nil
	}
	return true
}

// Error returns any error that may have been encountered.
func (it *instructionIterator) Error() error {
	return it.error
}

// PC returns the PC of the current instruction.
func (it *instructionIterator) PC() uint64 {
	return it.pc
}

// Op returns the opcode of the current instruction.
func (it *instructionIterator) Op() vm.OpCode {
	return it.op
}

// Arg returns the argument of the current instruction.
func (it *instructionIterator) Arg() []byte {
	return it.arg
}


func main() {
	fmt.Println("The EVM foot soldier")
  
  code := "608060405234801561001057600080fd5b5060fb8061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c806342e2d4d914602d575b600080fd5b605860048036036040811015604157600080fd5b506001600160a01b0381358116916020013516605a565b005b6040805163095ea7b360e01b81526001600160a01b0383811660048301526003602483015291519184169163095ea7b39160448082019260009290919082900301818387803b15801560ab57600080fd5b505af115801560be573d6000803e3d6000fd5b50505050505056fea265627a7a7231582074d2a91f880b7effa21aa471ea48b6bcf75305a0c1e9d86c0e0fda548a127b2464736f6c63430005110032"

	fmt.Printf("%v\n", code)

	script, err := hex.DecodeString(code)
	if err != nil {
    fmt.Println("Error decoding")
	}

	it := NewInstructionIterator(script)
	for it.Next() {
		if it.Arg() != nil && 0 < len(it.Arg()) {
			fmt.Printf("%05x: %v %#x\n", it.PC(), it.Op(), it.Arg())
		} else {
			fmt.Printf("%05x: %v\n", it.PC(), it.Op())
		}
	}
}
