package ui

import (
	"fmt"
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
	tbl := table.NewInstructionTable(contract.Instructions, config)

	flex := createFlex(contract)
	flex.AddItem(tbl.View, 0, 1, true)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	return UI{
		Application: app,
		Config:      config,
	}
}

func createFlex(contract eth.Contract) *tview.Flex {
	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)
	flex.SetBorder(true)

	if contract.IsMock {
		flex.SetTitle("[ Troy ]")
	} else {
		flex.SetTitle(fmt.Sprintf("[ %s ]", contract.Address))
	}

	return flex
}
