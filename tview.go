package main

import (
	"strings"

	"github.com/rivo/tview"
)

func TUI() {
	app := tview.NewApplication()
	pages := tview.NewPages()

	list := tview.NewList()
	list.AddItem("Create Post", "", 0, func() {
		Form(pages)
	}).
		AddItem("Read Post", "", 0, func() {
			Search_choice(pages, 'r')
		}).
		AddItem("Update Post", "", 0, func() {
			Search_choice(pages, 'u')
		}).
		AddItem("Delete Post", "", 0, func() {
			Search_choice(pages, 'd')
		}).
		AddItem("Exit", "", 0, func() { app.Stop() }).
		ShowSecondaryText(false).
		SetWrapAround(true).
		SetBorder(true).
		SetTitle("What do you want to do?")

	pages.AddAndSwitchToPage("list", list, true)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func Form(pages *tview.Pages) {
	title := tview.NewInputField().
		SetLabel("Title: ").
		SetFieldWidth(20)

	body := tview.NewTextArea().
		SetLabel("Body: ")

	form := tview.NewForm().
		AddFormItem(title).
		AddFormItem(body).
		AddButton("Create", func() {
			pages.SwitchToPage("list")
			// Create()
		}).
		AddButton("Cancel", func() {
			pages.SwitchToPage("list")
		})
	pages.AddAndSwitchToPage("form", form, true)
}

func Search_choice(pages *tview.Pages, r rune) {
	choice := tview.NewModal()
	choice.SetText("How do you want to search the post?\n").
		AddButtons([]string{"By Id (Exact Search)", "By Title (Similar Search)"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonIndex == 0 {
				Search_Id(pages, r)
			} else {
				Search_Title(pages, r)
			}
		})
	pages.AddAndSwitchToPage("search", choice, true)
}

func Search_Id(pages *tview.Pages, r rune) {
	form := tview.NewForm().
		AddInputField("Id: ", "", 10, tview.InputFieldInteger, nil).
		AddButton("Submit", func() {
			switch r {
			case 'r': // implement func
			case 'u':
			case 'd':
			}
		}).AddButton("Cancel", func() {
		pages.SwitchToPage("list")
	})
	pages.AddAndSwitchToPage("id", form, true)
}

func Search_Title(pages *tview.Pages, r rune) {
	posts := [][]string{
		{"Golang Basics", "Text for Basics"},
		{"Gin Framework", "Text for Framework"},
		{"Building REST APIs", "Text for API"},
		{"Concurrency in Go", "Text for concurrency"},
		{"Understanding Channels", "Text for Channels"},
		{"Go Routines vs Threads", "Text for threads"},
		{"Microservices with Go", "Text for Microservices"},
		{"Database Handling in Go", "Text for database"},
		{"Go Memory Management", "Text for memory management"},
	} // Take the posts from web

	list := tview.NewList()
	list.SetWrapAround(true).
		SetBorder(true).
		SetTitle("POSTS:-")

	update := func(text string) {
		list.Clear()
		for _, post := range posts {
			if strings.Contains(strings.ToLower(post[0]), strings.ToLower(text)) {
				list.AddItem(post[0], post[1], 0, func() {
					pages.SwitchToPage("list")
				}) // Change the function
			}
		}
	}

	search_bar := tview.NewForm().
		AddInputField("Title: ", "", 20, nil, func(text string) {
			update(text)
		}).AddButton("Cancel", func() { // The button is not showing yet
		pages.SwitchToPage("list")
	})

	update("")

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(search_bar, 3, 1, true).
		AddItem(list, 0, 1, false)
	pages.AddAndSwitchToPage("title", layout, true)
}
