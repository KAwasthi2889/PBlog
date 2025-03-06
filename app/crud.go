package main

import "errors"

func Create(post *Post) {
	if post.ID == 0 {
		db.Create(post)
		db.Last(post)
		Notify(post, "Create")
	} else {
		db.Save(post)
		Notify(post, "Update")
	}
}

func Read(id int) (*Post, error) {
	var post Post
	db.First(&post, id)
	if post.ID == 0 {
		return nil, errors.New("ID Dosen't Exist")
	}
	return &post, nil
}

func Update(id int) {
	post, err := Read(id)
	if err != nil {
		Notify(nil, err.Error())
	} else {
		Form(post.ID, post.Title, post.Body)
	}
}

func Delete(id int) {
	post, err := Read(id)
	if err != nil {
		Notify(nil, err.Error())
	} else {
		db.Delete(&Post{}, id)
		Notify(post, "Delete")
	}
}
