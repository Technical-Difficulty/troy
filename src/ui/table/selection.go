package table

import (
	"fmt"
	"troy/src/dasm"
)

type SelectionChange struct {
	Row          int
	PreviousText string
}

func (t *ContractTable) handleSelectionChanged(row int, column int) {
	// Don't process a row that is marked as not selectable
	if t.View.GetCell(row, column).NotSelectable {
		return
	}

	// Restore any previously made changes
	t.restoreChanges()

	ins := t.InstructionMap.GetInstruction(row)

	switch ins.OpCode.String() {
	case "JUMPI":
		t.handleJumpSelection(ins)
		break
	}
}

func (t *ContractTable) handleJumpSelection(ins dasm.Instruction) {
	if destPC, found := t.Analysis.JumpMap[ins.PC]; found {
		item := t.InstructionMap.GetFromPC(destPC)
		destCell := t.View.GetCell(item.Index, 0)

		t.storeChange(item.Index, destCell.Text)
		destCell.SetText(fmt.Sprintf("  [black:yellow:b][%05x][yellow:black:bl] JUMPDEST", item.Instruction.PC))
	}
}

func (t *ContractTable) storeChange(row int, text string) {
	t.changes = append(t.changes, SelectionChange{
		Row:          row,
		PreviousText: text,
	})
}

func (t *ContractTable) restoreChanges() {
	for _, c := range t.changes {
		cell := t.View.GetCell(c.Row, 0)
		cell.SetText(c.PreviousText)
	}
}
