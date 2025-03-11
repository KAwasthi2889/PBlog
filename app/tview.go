package main

import (
	"strconv"

	"github.com/rivo/tview"
)

var pages *tview.Pages

func TUI() {
	app := tview.NewApplication()
	pages = tview.NewPages()

	list := tview.NewList()
	list.
		AddItem("\n▶ Create Post", "", 0, func() { Form(0, "", "", "") }).
		AddItem("\n▶ Read Post", "", 0, func() { Search_choice('r') }).
		AddItem("\n▶ Update Post", "", 0, func() { Search_choice('u') }).
		AddItem("\n▶ Delete Post", "", 0, func() { Search_choice('d') }).
		AddItem("\n❌ Exit", "", 0, func() { app.Stop() }).
		ShowSecondaryText(false).
		SetWrapAround(true).
		SetBorder(true).
		SetTitle(" COMMANDS ").
		SetTitleAlign(tview.AlignCenter)

	pages.AddAndSwitchToPage("commands", list, true)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func Form(prev_id int, prev_category, prev_title, prev_body string) {
	var newPost Post
	work := " CREATE "
	if prev_id != 0 {
		newPost.ID = prev_id
		work = " UPDATE "
	}

	title := tview.NewInputField().
		SetLabel(" TITLE: ").
		SetText(prev_title).
		SetFieldWidth(40)

	body := tview.NewTextArea().
		SetText(prev_body, true).
		SetLabel(" BODY: ")

	categories := []string{"Technology", "Personal", "Sports", "Business", "Entertainment", "Fashion"}
	category := tview.NewDropDown().
		SetLabel(" CATEGORY: ").
		SetFieldWidth(20)
	for i, option := range categories {
		category.AddOption(option, func() {
			_, newPost.Category = category.GetCurrentOption()
		})
		if option == prev_category {
			category.SetCurrentOption(i)
		}
	}

	form := tview.NewForm()
	form.
		AddFormItem(title).
		AddFormItem(body).
		AddFormItem(category).
		AddButton("✔"+work, func() {
			newPost.Title = title.GetText()
			newPost.Body = body.GetText()
			Create(&newPost)
		}).
		AddButton("✖ CANCEL", func() { pages.SwitchToPage("commands") }).
		SetBorder(true).
		SetTitle(work + "POST").
		SetTitleAlign(tview.AlignCenter)

	pages.AddAndSwitchToPage("create", form, true)
}

func Notify(post *Post, work string) {
	var message string
	modal := tview.NewModal()

	if post != nil {
		work += "d"
		message = "Post " + work + " successfully!!\n"
		modal.AddButtons([]string{"📖 View " + work + " Post", "↩ Return"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonIndex == 0 {
					Reader(post)
				} else {
					pages.SwitchToPage("commands")
				}
			})
	} else {
		message = work
		modal.AddButtons([]string{"↩ Return"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				pages.SwitchToPage("commands")
			})
	}

	modal.SetText(message).SetTextColor(tview.Styles.PrimaryTextColor)
	pages.AddAndSwitchToPage("modal", modal, true)
}

func Reader(post *Post) {
	title := "TITLE: " + post.Title + " || ID: " + strconv.Itoa(post.ID) + " || CATEGORY: " + post.Category
	body := "\n\nBODY: " + post.Body
	content := title + body

	form := tview.NewForm()
	form.
		AddTextView("", content, 0, 0, false, true).
		AddButton("↩ Return", func() { pages.SwitchToPage("commands") }).
		SetBorder(true).
		SetTitle(" POST DETAILS ").
		SetTitleAlign(tview.AlignCenter)

	pages.AddAndSwitchToPage("reader", form, true)
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
		AddButton("Submit", func() { Select_Work(id, r) }).
		AddButton("Cancel", func() { pages.SwitchToPage("commands") })
	pages.AddAndSwitchToPage("id", form, true)
}

func Search_Title(r rune) {
	posts := []Post{}
	title := ""

	list := tview.NewList()
	list.SetWrapAround(true).
		SetBorder(true).
		SetTitle("POSTS")

	updateList := func() {
		list.Clear()
		for _, post := range posts {
			list.AddItem(post.Title, "", 0, func() {
				Select_Work(post.ID, r)
			})
		}
	}

	last_id := Multiple(&posts, title, 0)
	updateList()

	search_bar := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(tview.NewInputField().
			SetLabel("Title:  ").
			SetChangedFunc(func(text string) {
				title = text
				last_id = Multiple(&posts, title, 0)
				updateList()
			}), 0, 17, true).
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(tview.NewButton("Cancel").
			SetSelectedFunc(func() {
				pages.SwitchToPage("commands")
			}), 0, 2, true)

	button := tview.NewButton("Next").
		SetSelectedFunc(func() {
			last_id = Multiple(&posts, title, last_id)
			if len(posts) == 0 {
				last_id = Multiple(&posts, title, 0)
			}
			updateList()
		})

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(search_bar, 1, 0, true).
		AddItem(list, 0, 1, true).
		AddItem(button, 1, 0, true)

	pages.AddAndSwitchToPage("title", layout, true)
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

func Search_choice(r rune) {
	choice := tview.NewModal()
	choice.SetText("How do you want to search the post?\n").
		AddButtons([]string{"By Id (Exact Search)", "By Title (Similar Search)"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonIndex == 0 {
				Search_Id(r)
			} else {
				Search_Title(r)
			}
		})
	pages.AddAndSwitchToPage("search", choice, true)
}
