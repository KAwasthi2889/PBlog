package main

import "errors"

func Create(post *Post) {
	if post.ID == 0 {
		var duplicate Post
		db.First(&duplicate, "Title = ?", post.Title)
		if duplicate.ID == 0 {
			db.Create(post)
			db.Last(post)
			Notify(post, "Create")
		} else {
			Notify(nil, "Post with same title already exists, Please try again!!")
		}

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

		Form(post.ID, post.Category, post.Title, post.Body)
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

func Multiple(posts *[]Post, Title string, last_id int) int {
	simmilar := "%" + Title + "%"
	db.Where("title LIKE ? AND id > ?", simmilar, last_id).Limit(5).Find(posts)
	last_id = 0
	for _, post := range *posts {
		last_id = post.ID
	}
	return last_id
}
