package ui

import (
	"github.com/rivo/tview"
	"troy/src"
	"troy/src/eth"
	"troy/src/ui/table"
)

type UI struct {
	Application *tview.Application
	Config      src.Config
}

func Start(contract eth.Contract, config src.Config) UI {
	app := tview.NewApplication()
	tbl := table.NewInstructionTable(contract.Instructions)

	flex := createFlex()
	flex.AddItem(tbl.View, 0, 1, true)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	return UI{
		Application: app,
		Config:      config,
	}
}

func createFlex() *tview.Flex {
	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)
	flex.SetTitle("[ Troy ]")
	flex.SetBorder(true)

	return flex
}
