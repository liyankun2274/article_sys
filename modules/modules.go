package modules

import "time"

type User struct {
	ID       int    `json:"id"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Status   int8   `json:"status"`
	CreateAt time.Time `json:"status"`
}

type Column struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status int8   `json:"status"`
	CreateAt time.Time `json:"status"`
	UpdateAt time.Time `json:"status"`
	DeleteAt *time.Time `json:"status"`
}

type Article struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	ColumnID int    `json:"column_id"`
	UserId   int    `json:"user_id"`
	Author   string `json:"author"`
	Status   int8   `json:"status"`
	CreateAt time.Time `json:"status"`
	UpdateAt time.Time `json:"status"`
	DeleteAt *time.Time `json:"status"`
}
