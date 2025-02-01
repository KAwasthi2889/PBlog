package main

import (
	"fmt"

	"github.com/rivo/tview"
)

func TUI() {
	app := tview.NewApplication() // Can I remove letters and make them look like >

	list := tview.NewList().
		AddItem("Create Post", "", 'c', func() {
			app.Stop()
			OpenForm()
		}).
		AddItem("Search Post", "", 's', func() {
			app.Stop()
			SearchForm('s')
		}).
		AddItem("Update Post", "", 'u', func() {
			app.Stop()
			SearchForm('u')
		}).
		AddItem("Delete Post", "", 'd', func() {
			app.Stop()
			SearchForm('d')
		}).
		AddItem("Exit", "", 'e', func() { app.Stop() }).ShowSecondaryText(false).SetWrapAround(true)

	list.SetBorder(true).SetTitle("What do you want to do?")

	if err := app.SetRoot(list, true).EnableMouse(true).Run(); err != nil { // Does Enable Mouse Work?
		panic(err)
	}
}

func OpenForm() {
	app := tview.NewApplication()

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

	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func SearchForm(r rune) {
	app := tview.NewApplication()

	app.Run()
	switch r {
	case 'u':
	case 'd':
	}
}
