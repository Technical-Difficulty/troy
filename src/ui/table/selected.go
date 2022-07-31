package table

import "troy/src/dasm"

func (t *ContractTable) handleSetSelected(row int, column int) {
	ins := t.InstructionMap.GetInstruction(row)

	switch ins.OpCode.String() {
	case "JUMPI":
		t.handleJumpSelected(ins)
		break
	}
}

func (t *ContractTable) handleJumpSelected(ins dasm.Instruction) {
	if pc, found := t.Analysis.JumpMap[ins.PC]; found {
		destination := t.InstructionMap.GetIndex(pc)

		t.View.Select(destination, 0)
	}
}
