package main

type Post struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Title    string `json:"title"`
	Body     string `json:"content"`
	Category string `json:"category"`
}
