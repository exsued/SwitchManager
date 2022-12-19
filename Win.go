package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func ScanWinExec(a fyne.App, usedIface string, prevWin fyne.Window) {
	w := a.NewWindow("Switch manager")
	w.Resize(fyne.NewSize(400, 100))
	label := widget.NewLabel("Scanning...")
	progBar := widget.NewProgressBar()
	c := container.NewVBox(
		label,
		progBar,
	)
	w.SetContent(c)
	w.SetMaster()
	prevWin.Close()
	w.Show()
}

func SelectWinExec(a fyne.App, ifaces []string) {
	w := a.NewWindow("Select iface")
	w.Resize(fyne.NewSize(400, 200))
	var usedIface string
	selectIface := widget.NewSelect(ifaces, func(s string) {
		usedIface = s
	})
	button := widget.NewButton("Continue", func() {
		w.Hide()
		ScanWinExec(a, usedIface, w)
	})
	c := container.NewVBox(
		selectIface,
		button,
	)
	w.SetContent(c)
	w.ShowAndRun()
}
