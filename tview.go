package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func TUI() {
	app := tview.NewApplication()
	list := tview.NewList().
		AddItem("Create Post", "", 'c', func() {
			OpenForm(app)
			defer app.Stop()
		}).
		AddItem("Search Post", "", 's', func() {
			Search()
			defer app.Stop()
		}).
		AddItem("Update Post", "", 'u', func() {
			Update()
			defer app.Stop()
		}).
		AddItem("Delete Post", "", 'd', func() {
			Delete()
			defer app.Stop()
		}).
		AddItem("Exit", "", 'e', func() { app.Stop() }).ShowSecondaryText(false)

	frame := tview.NewFrame(list).
		AddText("What do you want to do?", true, tview.AlignLeft, tcell.ColorWhiteSmoke)

	if err := app.SetRoot(frame, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}

func OpenForm(app *tview.Application) {
	title := tview.NewInputField().
		SetLabel("Title: ").
		SetFieldWidth(20)

	body := tview.NewTextArea().
		SetLabel("Body: ")

	form := tview.NewForm().
		AddFormItem(title).
		AddFormItem(body).
		AddButton("Create", func() {
			Create(title.GetText(), body.GetText())
		}).
		AddButton("Cancel", func() {
			fmt.Println("Cancel")
			app.Stop()
		})

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 1, true)

	if err := app.SetRoot(layout, true).SetFocus(form).Run(); err != nil {
		panic(err)
	}
}
