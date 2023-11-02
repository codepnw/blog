package models

type Blog struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Title string `json:"title" gorm:"not null;column:title;size:255"`
	Post  string `json:"post" gorm:"not null;column:post;size:255"`
}
