package ui

import (
	"github.com/rivo/tview"
	"troy/src/dasm"
	"troy/src/ui/table"
)

type UI struct {
	Application *tview.Application
}

func Start(instructions []dasm.Instruction) {
	app := tview.NewApplication()
	tbl := table.NewInstructionTable(instructions)

	flex := createFlex()
	flex.AddItem(tbl.View, 0, 1, true)

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
