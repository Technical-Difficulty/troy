package table

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"troy/src"
	"troy/src/dasm"
	"troy/src/dasm/enum"
)

type ContractTable struct {
	View         *tview.Table
	Instructions dasm.InstructionSet
	Analysis     enum.ContractAnalysis
	row          int
	config       src.Config
}

func NewContractTable(analysis enum.ContractAnalysis, config src.Config) ContractTable {
	table := ContractTable{
		View:         tview.NewTable(),
		Instructions: analysis.Contract.Instructions,
		Analysis:     analysis,
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

func (t *ContractTable) init() {
	colors := t.config.Colors.Table

	t.View.SetSelectable(true, false)
	t.View.SetBackgroundColor(t.getColor(colors.Default.Background))
	t.View.SetSelectedStyle(tcell.StyleDefault.
		Foreground(t.getColor(colors.Selected.Foreground)).
		Background(t.getColor(colors.Selected.Background)))
}

func (t *ContractTable) addRowCell(cell *tview.TableCell, selectable bool) {
	t.View.SetCell(t.row, 0, cell.SetSelectable(selectable))
	t.row++
}

func (t *ContractTable) addRow(text string, selectable bool) {
	cell := tview.NewTableCell(text).SetSelectable(selectable)
	t.View.SetCell(t.row, 0, cell)
	t.row++
}

func (t *ContractTable) addBoldRow(text string, selectable bool) {
	cell := tview.NewTableCell(text).SetSelectable(selectable).SetStyle(tcell.StyleDefault.Bold(true))
	t.View.SetCell(t.row, 0, cell)
	t.row++
}

func (t *ContractTable) addEmptyRow(selectable bool) {
	t.View.SetCell(t.row, 0, tview.NewTableCell("").SetSelectable(selectable))
	t.row++
}

func (t *ContractTable) addInstruction(ins dasm.Instruction) {
	cell := tview.NewTableCell(t.getInstructionOutput(ins)).
		SetExpansion(1).
		SetAlign(tview.AlignLeft)

	t.instructionPrefix(ins)
	t.addRowCell(cell, true)
}
