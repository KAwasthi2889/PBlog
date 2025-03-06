package main

func Create(post Post) {
	if post.ID == -1 {
		db.Create(&post)
		Notify(&post, "Create")
	} else {
		db.Save(&post)
		Notify(&post, "Update")
	}
}

func Read(id int) (*Post, error) {
	var count int64
	err := db.Model(&Post{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return nil, err
	}
	var post Post
	db.First(&post, id)
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
