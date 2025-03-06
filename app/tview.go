package main

import (
	"strconv"

	// "strings"

	"github.com/rivo/tview"
)

var pages *tview.Pages

func TUI() {
	app := tview.NewApplication()
	pages = tview.NewPages()

	list := tview.NewList()
	list.AddItem("Create Post", "", 0, func() {
		Form(-1, "", "")
	}).
		AddItem("Read Post", "", 0, func() {
			Search_choice('r')
		}).
		AddItem("Update Post", "", 0, func() {
			Search_choice('u')
		}).
		AddItem("Delete Post", "", 0, func() {
			Search_choice('d')
		}).
		AddItem("Exit", "", 0, func() { app.Stop() }).
		ShowSecondaryText(false).
		SetWrapAround(true).
		SetBorder(true).
		SetTitle("What do you want to do?")

	pages.AddAndSwitchToPage("commands", list, true)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func Form(prev_id int, prev_title, prev_body string) {
	var newPost Post
	newPost.ID = prev_id

	title := tview.NewInputField().
		SetLabel("Title: ").
		SetText(prev_title).
		SetFieldWidth(20)

	body := tview.NewTextArea().
		SetText(prev_body, true).
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
			newPost.Title = title.GetText()
			newPost.Body = body.GetText()
			Create(&newPost)
		}).
		AddButton("Cancel", func() {
			pages.SwitchToPage("commands")
		})
	pages.AddAndSwitchToPage("create", form, true)
}

func Search_choice(r rune) {
	choice := tview.NewModal()
	choice.SetText("How do you want to search the post?\n").
		AddButtons([]string{"By Id (Exact Search)", "By Title (Similar Search)"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonIndex == 0 {
				Search_Id(r)
			} else {
				// Search_Title(r)
			}
		})
	pages.AddAndSwitchToPage("search", choice, true)
}

func Search_Id(r rune) {
	var id int
	form := tview.NewForm().
		AddInputField("Id: ", "", 10, tview.InputFieldInteger, func(text string) {
			num, err := strconv.Atoi(text)
			if err != nil {
				Notify(nil, err.Error())
			}
			id = num
		}).
		AddButton("Submit", func() {
			Select_Work(id, r)
		}).
		AddButton("Cancel", func() {
			pages.SwitchToPage("commands")
		})
	pages.AddAndSwitchToPage("id", form, true)
}

func Select_Work(id int, r rune) {
	switch r {
	case 'r':
		post, err := Read(id)
		if err != nil {
			Notify(nil, err.Error())
		} else {
			Reader(post)
		}
	case 'u':
		Update(id)
	case 'd':
		Delete(id)
	}
}

func Notify(post *Post, work string) {
	var message string
	modal := tview.NewModal()
	if post != nil {
		work += "d"
		message = "Post " + work + " Successfully!!!\n"
		modal.AddButtons([]string{"View " + work + " Post", "Return"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonIndex == 0 {
					Reader(post)
				} else {
					pages.SwitchToPage("commands")
				}
			})
	} else {
		message = work
		modal.AddButtons([]string{"Return"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				pages.SwitchToPage("commands")
			})
	}

	modal.SetText(message)
	pages.AddAndSwitchToPage("modal", modal, true)
}

func Reader(post *Post) {
	title := "Title: " + post.Title + "#" + strconv.Itoa(post.ID)
	body := "\n\n" + post.Body
	content := title + body
	form := tview.NewForm().
		AddTextView("Content", content, 0, 0, false, true).
		AddButton("Return", func() {
			pages.SwitchToPage("commands")
		})
	pages.AddAndSwitchToPage("reader", form, true)
}

// func Search_Title(r rune) {
// 	posts := [][]string{
// 		{"Golang Basics", "Text for Basics"},
// 		{"Gin Framework", "Text for Framework"},
// 		{"Building REST APIs", "Text for API"},
// 		{"Concurrency in Go", "Text for concurrency"},
// 		{"Understanding Channels", "Text for Channels"},
// 		{"Go Routines vs Threads", "Text for threads"},
// 		{"Microservices with Go", "Text for Microservices"},
// 		{"Database Handling in Go", "Text for database"},
// 		{"Go Memory Management", "Text for memory management"},
// 		} // Take the posts from web

// 		list := tview.NewList()
// 		list.SetWrapAround(true).
// 		SetBorder(true).
// 		SetTitle("POSTS")

// 	update := func(text string) {
// 		list.Clear()
// 		for _, post := range posts {
// 			if strings.Contains(strings.ToLower(post[0]), strings.ToLower(text)) {
// 				list.AddItem(post[0], post[1], 0, func() {
// 					Select_Work(r)
// 				})
// 			}
// 		}
// 	}
// 		search_bar := tview.NewFlex().
// 		SetDirection(tview.FlexColumn).
// 		AddItem(tview.NewInputField().
// 		SetLabel("Title:  ").
// 		SetChangedFunc(func(text string) {
// 			update(text)
// 		}),
// 		0, 17, true).
// 		AddItem(tview.NewBox(), 0, 1, false).
// 		AddItem(tview.NewButton("Cancel").
// 		SetSelectedFunc(func() {
// 			pages.SwitchToPage("list")
// 		}),
// 		0, 2, true)

// 		update("")

// 		layout := tview.NewFlex().
// 		SetDirection(tview.FlexRow).
// 		AddItem(search_bar, 1, 0, true).
// 		AddItem(list, 0, 1, false)

// 		pages.AddAndSwitchToPage("title", layout, true)
// 		}
