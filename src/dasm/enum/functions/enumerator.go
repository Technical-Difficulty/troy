package functions

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
	"troy/src/dasm"
	"troy/src/util"
)

type FunctionEnum struct {
	instructions dasm.InstructionSet
	signatures   *map[string][]string
}

func NewFunctionEnum(instructions dasm.InstructionSet) FunctionEnum {
	sigs, err := loadSignaturesFromFile("data/signatures.json")
	if err != nil {
		log.WithField("Error", err).Fatal("Failed to load function signatures from file")
	}

	return FunctionEnum{
		instructions: instructions,
		signatures:   sigs,
	}
}

func loadSignaturesFromFile(path string) (*map[string][]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	signatures := make(map[string][]string)
	err = json.Unmarshal(data, &signatures)
	if err != nil {
		return nil, err
	}

	return &signatures, nil
}

// Function sigantures live between:
// PUSH1 0x04
// CALLDATASIZE
// ...
// CALLDATALOAD
// ...
// PUSH4 <FUNC_SIG>
// ...
// JUMPDEST
func (e *FunctionEnum) EnumerateSignatures() map[uint64][]string {
	var scan bool
	var sig Signature
	var mapping = make(map[uint64][]string)

	for idx, ins := range e.instructions.Array() {
		if ins.OpCode.String() == "CALLDATALOAD" {
			scan = true
		}

		if scan && ins.OpCode.String() == "PUSH4" {
			sig = NewSignature(ins)
		}

		if scan && ins.OpCode.String() == "JUMPI" {
			found, err := e.instructions.ReverseFindInstructionWithPrefix(idx, "PUSH")
			if err != nil {
				log.WithField("Error", err).Fatal("Couldn't find preceding PUSH instruction")
			}

			// As we are on the JUMPI instruction, we need to find the PC of the
			// operand instruction/location
			destinationPC, err := util.BytesToUInt64(found.Operand)
			mapping[destinationPC] = sig.Lookup(e.signatures)
		}

		if scan && ins.OpCode.String() == "JUMPDEST" {
			break
		}
	}

	return mapping
}
