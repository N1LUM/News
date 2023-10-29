package models

import (
	"time"
)

type Paper struct {
	ID         string     `gorm:"primaryKey" json:"id" form:"id"`
	TitleImage string     `json:"titleImage" form:"titleImage"`
	Title      string     `json:"title" form:"title"`
	Subtitle   string     `json:"subtitle" form:"subtitle"`
	Text       string     `json:"text" form:"text"`
	Image      *string    `json:"image" form:"image"`
	Author     Author     `json:"author" form:"author"`
	CreatedAt  time.Time  `json:"createdAt" form:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt" form:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt" form:"deletedAt"`
}

type Author struct {
	ID        string     `gorm:"primaryKey" json:"id" form:"id"`
	Username  string     `json:"username" form:"username"`
	Name      string     `json:"name" form:"name"`
	Lastname  string     `json:"lastname" form:"lastname"`
	Sex       string     `json:"sex" form:"sex"`
	Age       int        `json:"age" form:"age"`
	BornDate  time.Time  `json:"bornDate" form:"bornDate"`
	Country   string     `json:"country" form:"country"`
	City      string     `json:"city" form:"city"`
	CreatedAt time.Time  `json:"createdAt" form:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt" form:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" form:"deletedAt"`
}

type PaperInput struct {
	TitleImage string  `json:"titleImage" form:"titleImage"`
	Title      string  `json:"title"form:"title" `
	Subtitle   string  `json:"subtitle"form:"subtitle" `
	Text       string  `json:"text" form:"text"`
	Image      *string `json:"image" form:"image"`
	Author     Author  `json:"author" form:"author"`
}

type AuthorCreateInput struct {
	Username string `json:"username" form:"username"`
	Name     string `json:"name" form:"name"`
	Lastname string `json:"lastname" form:"lastname"`
	Sex      string `json:"sex" form:"sex"`
	Age      int    `json:"age" form:"age"`
	BornDate string `json:"bornDate" form:"bornDate"`
	Country  string `json:"country" form:"country"`
	City     string `json:"city" form:"city"`
}
type AuthorUpdateInput struct {
	Username *string `json:"username" form:"username"`
	Name     *string `json:"name" form:"name"`
	Lastname *string `json:"lastname" form:"lastname"`
	Sex      *string `json:"sex" form:"sex"`
	Age      *int    `json:"age" form:"age"`
	BornDate *string `json:"bornDate" form:"bornDate"`
	Country  *string `json:"country" form:"country"`
	City     *string `json:"city" form:"city"`
}
