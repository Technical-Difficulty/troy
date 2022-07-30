package table

import (
	"fmt"
	"strings"
	"troy/src/dasm"
)

func (t *ContractTable) instructionPrefix(ins dasm.Instruction) {
	switch ins.OpCode.String() {
	case "JUMPDEST":
		t.outputInstructionBlockHeader(ins)
		break
	}
}

func (t *ContractTable) outputInstructionBlockHeader(ins dasm.Instruction) {
	t.addRow("", false)
	t.addBoldRow(fmt.Sprintf(":loc_0x%04x", ins.PC), false)

	if sigs, found := t.Analysis.SignatureMap[ins.PC]; found {
		t.outputFunctionSignature(sigs)
	}
}

// todo: If we detect a function but don't detect the signature, inform the user
func (t *ContractTable) outputFunctionSignature(sigs []string) {
	signature := sigs[0]
	possibleSignatures := strings.Join(sigs, ", ")

	t.addRow("--------------------------------------------------------------", false)
	t.addRow(fmt.Sprintf("  function %s", signature), false)
	t.addRow(fmt.Sprintf("  possible signatures: [%s]", possibleSignatures), false)
	t.addRow("--------------------------------------------------------------", false)
}
