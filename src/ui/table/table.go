package table

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"troy/src"
	"troy/src/dasm"
	"troy/src/dasm/enum"
)

type InstructionTable struct {
	View         *tview.Table
	Instructions dasm.InstructionSet
	Analysis     enum.ContractEnum
	row          int
	config       src.Config
}

func NewInstructionTable(enum enum.ContractEnum, config src.Config) InstructionTable {
	table := InstructionTable{
		View:         tview.NewTable(),
		Instructions: enum.Contract.Instructions,
		Analysis:     enum,
		row:          0,
		config:       config,
	}

	// Any special configuration can be performed here
	table.init()

	for _, ins := range table.Instructions.Array() {
		table.addInstruction(ins)
	}

	return table
}

func (t *InstructionTable) init() {
	colors := t.config.Colors.Table

	t.View.SetSelectable(true, false)
	t.View.SetBackgroundColor(t.getColor(colors.Default.Background))
	t.View.SetSelectedStyle(tcell.StyleDefault.
		Foreground(t.getColor(colors.Selected.Foreground)).
		Background(t.getColor(colors.Selected.Background)))
}

func (t *InstructionTable) addColumn(cell *tview.TableCell, selectable bool) {
	t.View.SetCell(t.row, 0, cell.SetSelectable(selectable))
	t.row++
}

func (t *InstructionTable) addRow(text string, selectable bool) {
	cell := tview.NewTableCell(text).SetSelectable(selectable)
	t.View.SetCell(t.row, 0, cell)
	t.row++
}

func (t *InstructionTable) addBoldRow(text string, selectable bool) {
	cell := tview.NewTableCell(text).SetSelectable(selectable).SetStyle(tcell.StyleDefault.Bold(true))
	t.View.SetCell(t.row, 0, cell)
	t.row++
}

func (t *InstructionTable) addEmptyRow(selectable bool) {
	t.View.SetCell(t.row, 0, tview.NewTableCell("").SetSelectable(selectable))
	t.row++
}

func (t *InstructionTable) addInstruction(ins dasm.Instruction) {
	cell := tview.NewTableCell(t.getInstructionOutput(ins)).
		SetExpansion(1).
		SetAlign(tview.AlignLeft)

	t.instructionPrefix(ins)
	t.addColumn(cell, true)
}
