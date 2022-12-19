package main

import (
	"fyne.io/fyne/app"
)

func main() {
	a := app.New()
	ifaceNames, _ := getIfaceNames()
	SelectWinExec(a, ifaceNames)
}
