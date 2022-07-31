package table

import (
	"fmt"
	"strings"
	"troy/src/dasm"
)

func (t *ContractTable) instructionPrefix(ins dasm.Instruction) {
	if ins.PC == 0 {
		t.outputInstructionBlockHeader(ins)
	}

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

func (t *ContractTable) outputFunctionSignature(sigs []string) {
	var signature string
	var possibleSignatures string

	if sigs == nil {
		signature = ""
		possibleSignatures = "none found"
	} else {
		signature = sigs[0]
		possibleSignatures = strings.Join(sigs, ", ")
	}

	c := t.config.Colors.Table.FunctionSignature
	tag := t.getDefaultColorTag()

	t.addRow(fmt.Sprintf("%s--------------------------------------------------------------%s", c.Divider, tag), false)
	t.addRow(fmt.Sprintf("%s  function %s%s", c.Function, signature, tag), false)
	t.addRow(fmt.Sprintf("%s  possible signatures: [%s]%s", c.PossibleSignatures, possibleSignatures, tag), false)
	t.addRow(fmt.Sprintf("%s--------------------------------------------------------------%s", c.Divider, tag), false)
}
