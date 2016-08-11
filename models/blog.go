package models

import (
	"time"
)

type Blog struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

type Blogs []Blog
