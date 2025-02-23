package main

type Post struct {
	Id       int    `json:"Id"`
	Title    string `json:"Title"`
	Body     string `json:"Content"`
	Category string `json:"Category"`
}

var DB *grom.DB
