package ui

import (
	"fmt"
	"github.com/rivo/tview"
	"troy/src/dasm"
)

type InstructionTable struct {
	View         *tview.Table
	Instructions []dasm.Instruction
}

func NewInstructionTable(instructions []dasm.Instruction) InstructionTable {
	table := tview.NewTable()
	table.SetSelectable(true, false)

	for i, ins := range instructions {
		var output string

		if ins.Operand != nil {
			output = fmt.Sprintf("[%05x] %v %#x\n", ins.PC, ins.OpCode, ins.Operand)
		} else {
			output = fmt.Sprintf("[%05x] %v\n", ins.PC, ins.OpCode)
		}

		table.SetCell(i, 0, tview.NewTableCell(output).SetExpansion(1).SetAlign(tview.AlignLeft))
	}

	return InstructionTable{
		View:         table,
		Instructions: instructions,
	}
}
