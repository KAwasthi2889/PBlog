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
	var newPost Post

	title := tview.NewInputField().
		SetLabel("Title: ").
		SetFieldWidth(20)

	body := tview.NewTextArea().
		SetLabel("Body: ")

	categories := []string{
		"Technology", "Personal", "Sports", "Business", "Entertainment", "Fashion",
	}

	category := tview.NewDropDown()

	for _, option := range categories {
		category.AddOption(option, func() {
			newPost.Category = option
		})
	}

	form := tview.NewForm().
		AddFormItem(title).
		AddFormItem(body).
		AddFormItem(category).
		AddButton("Create", func() {
			newPost.Body = body.GetText()
			newPost.Title = title.GetText()
			Create(newPost)
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
			Select_Work(pages, r)
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
		SetTitle("POSTS")

	update := func(text string) {
		list.Clear()
		for _, post := range posts {
			if strings.Contains(strings.ToLower(post[0]), strings.ToLower(text)) {
				list.AddItem(post[0], post[1], 0, func() {
					Select_Work(pages, r)
				})
			}
		}
	}

	search_bar := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(tview.NewInputField().
			SetLabel("Title:  ").
			SetChangedFunc(func(text string) {
				update(text)
			}),
			0, 17, true).
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(tview.NewButton("Cancel").
			SetSelectedFunc(func() {
				pages.SwitchToPage("list")
			}),
			0, 2, true)

	update("")

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(search_bar, 1, 0, true).
		AddItem(list, 0, 1, false)

	pages.AddAndSwitchToPage("title", layout, true)
}

func Select_Work(pages *tview.Pages, r rune) {
	switch r {
	case 'r':
		// ReadOnly(pages)
	case 'u':
		// ReadWrite(pages)
	case 'd':
		Delete()
		pages.SwitchToPage("list")
	}
}
