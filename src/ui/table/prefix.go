package table

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"troy/src/dasm"
)

func (t *InstructionTable) instructionPrefix(ins dasm.Instruction) {
	switch ins.OpCode.String() {
	case "JUMPDEST":
		t.handleJUMPDEST(ins)
		break
	}
}

func (t *InstructionTable) handleJUMPDEST(ins dasm.Instruction) {
	loc := fmt.Sprintf(":loc_0x%04x", ins.PC)
	s := tcell.StyleDefault.Bold(true)
	cell := tview.NewTableCell(loc).SetStyle(s)

	t.addColumn(tview.NewTableCell(""), false)
	t.addColumn(cell, false)
}
