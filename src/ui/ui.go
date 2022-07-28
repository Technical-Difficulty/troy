package ui

import (
	"github.com/rivo/tview"
	"troy/src/dasm"
)

type UI struct {
	Application *tview.Application
}

func Start(instructions []dasm.Instruction) {
	app := tview.NewApplication()
	table := NewInstructionTable(instructions)

	flex := createFlex()
	flex.AddItem(table.View, 0, 1, true)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func createFlex() *tview.Flex {
	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)
	flex.SetTitle("[ Troy ]")
	flex.SetBorder(true)

	return flex
}
