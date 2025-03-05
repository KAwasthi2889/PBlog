package main

func Create(post Post) {
	db.Save(&post)
	Modal("Saved")
}

func Read(id int) (Post, error) {
	var post Post
	result := db.First(&post, id)
	if result.Error != nil {
		return Post{}, result.Error
	}
	return post, nil
}

func Update(id int) {
	post, err := Read(id)
	if err != nil {
		Warn("Invalid ID!!!")
	}
	Form(post.Title, post.Body)
}

func Delete(id int) {
	db.Delete(&Post{}, id)
	Modal("Deleted")
}
