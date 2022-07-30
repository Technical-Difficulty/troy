package table

import (
	"github.com/rivo/tview"
	"troy/src"
	"troy/src/dasm"
)

type InstructionTable struct {
	View         *tview.Table
	Instructions []dasm.Instruction
	row          int
	config       src.Config
}

func NewInstructionTable(instructions []dasm.Instruction, config src.Config) InstructionTable {
	table := InstructionTable{
		View:         tview.NewTable(),
		Instructions: instructions,
		row:          0,
		config:       config,
	}

	// Any special configuration can be performed here
	table.init()

	for _, ins := range instructions {
		table.addInstruction(ins)
	}

	return table
}

func (t *InstructionTable) init() {
	t.View.SetSelectable(true, false)
}

func (t *InstructionTable) addColumn(cell *tview.TableCell, selectable bool) {
	t.View.SetCell(t.row, 0, cell.SetSelectable(selectable))
	t.row++
}

func (t *InstructionTable) addInstruction(ins dasm.Instruction) {
	cell := tview.NewTableCell(t.getInstructionOutput(ins)).
		SetExpansion(1).
		SetAlign(tview.AlignLeft)

	t.instructionPrefix(ins)
	t.addColumn(cell, true)
}
