package main

import "fmt"

func Create(title, body string) {
	fmt.Println("Create Post", title, body)
}

func Update() {
	fmt.Println("Update Post")
}

func Delete() {
	fmt.Println("Delete Post")
}
