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
	table := InstructionTable{
		View:         tview.NewTable(),
		Instructions: instructions,
	}

	table.init()

	for i, ins := range instructions {
		table.addColumn(i, ins)
	}

	return table
}

func (t *InstructionTable) init() {
	t.View.SetSelectable(true, false)
}

func (t *InstructionTable) addColumn(row int, ins dasm.Instruction) {
	var output string

	if ins.Operand != nil {
		output = fmt.Sprintf("[%05x] %v %#x\n", ins.PC, ins.OpCode, ins.Operand)
	} else {
		output = fmt.Sprintf("[%05x] %v\n", ins.PC, ins.OpCode)
	}

	t.View.SetCell(row, 0, tview.NewTableCell(output).SetExpansion(1).SetAlign(tview.AlignLeft))
}
