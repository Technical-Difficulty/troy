package ui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"troy/src/dasm"
)

type InstructionTable struct {
	View         *tview.Table
	Instructions []dasm.Instruction
	row          int
}

func NewInstructionTable(instructions []dasm.Instruction) InstructionTable {
	table := InstructionTable{
		View:         tview.NewTable(),
		Instructions: instructions,
		row:          0,
	}

	// Any special configuration
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
	var output string

	if ins.Operand != nil {
		output = fmt.Sprintf("  [%05x] %v %#x\n", ins.PC, ins.OpCode, ins.Operand)
	} else {
		output = fmt.Sprintf("  [%05x] %v\n", ins.PC, ins.OpCode)
	}

	cell := tview.NewTableCell(output).
		SetExpansion(1).
		SetAlign(tview.AlignLeft)

	t.instructionPrefix(ins)
	t.addColumn(cell, true)
}

func (t *InstructionTable) instructionPrefix(ins dasm.Instruction) {
	switch ins.OpCode.String() {
	case "JUMPDEST":
		t.JUMPDEST(ins)
		break
	}
}

func (t *InstructionTable) JUMPDEST(ins dasm.Instruction) {
	loc := fmt.Sprintf(":loc_0x%03x", ins.PC)
	s := tcell.StyleDefault.Bold(true)
	cell := tview.NewTableCell(loc).SetStyle(s)

	t.addColumn(tview.NewTableCell(""), false)
	t.addColumn(cell, false)
}
