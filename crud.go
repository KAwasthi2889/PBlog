package main

func Create(post Post) {
	db.Create(post)
	Modal("Post Created Succesfully!!!")
}

func Read() {
	//Use GET to read data
}

func Update() {
	Read()
	// Use PUT to set data
}

func Delete() {
	// Implement Delete
}

func Modal(message string) {

}
