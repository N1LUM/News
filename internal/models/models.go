package models

import (
	"mime/multipart"
	"time"
)

type Paper struct {
	ID            string     `gorm:"primaryKey" json:"id" form:"id"`
	TitleImageURL string     `json:"titleImageURL" form:"titleImageURL"`
	Title         string     `json:"title" form:"title"`
	Subtitle      string     `json:"subtitle" form:"subtitle"`
	Content       string     `json:"text" form:"text"`
	AuthorID      string     `json:"authorID" form:"authorID"`
	Author        Author     `json:"author" form:"author"`
	CreatedAt     time.Time  `json:"createdAt" form:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt" form:"updatedAt"`
	DeletedAt     *time.Time `json:"deletedAt" form:"deletedAt"`
}

type Author struct {
	ID        string     `gorm:"primaryKey" json:"id" form:"id"`
	Username  string     `json:"username" form:"username"`
	Name      string     `json:"name" form:"name"`
	Lastname  string     `json:"lastname" form:"lastname"`
	Sex       string     `json:"sex" form:"sex"`
	BornDate  time.Time  `json:"bornDate" form:"bornDate"`
	Country   string     `json:"country" form:"country"`
	City      string     `json:"city" form:"city"`
	CreatedAt time.Time  `json:"createdAt" form:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt" form:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" form:"deletedAt"`
}

type AuthorCreateInput struct {
	Username string `json:"username" form:"username"`
	Name     string `json:"name" form:"name"`
	Lastname string `json:"lastname" form:"lastname"`
	Sex      string `json:"sex" form:"sex"`
	BornDate string `json:"bornDate" form:"bornDate"`
	Country  string `json:"country" form:"country"`
	City     string `json:"city" form:"city"`
}
type AuthorUpdateInput struct {
	Username *string `json:"username" form:"username"`
	Name     *string `json:"name" form:"name"`
	Lastname *string `json:"lastname" form:"lastname"`
	Sex      *string `json:"sex" form:"sex"`
	BornDate *string `json:"bornDate" form:"bornDate"`
	Country  *string `json:"country" form:"country"`
	City     *string `json:"city" form:"city"`
}

type PaperCreateInput struct {
	TitleImage    multipart.FileHeader `json:"titleImage" form:"titleImage"`
	TitleImageURL string               `json:"titleImageURL" form:"titleImageURL"`
	Title         string               `json:"title" form:"title"`
	Subtitle      string               `json:"subtitle" form:"subtitle"`
	Content       string               `json:"text" form:"text"`
	AuthorID      string               `json:"authorID" form:"authorID"`
}
type PaperUpdateInput struct {
	TitleImage    *multipart.FileHeader `json:"titleImage" form:"titleImage"`
	TitleImageURL *string               `json:"titleImageURL" form:"titleImageURL"`
	Title         *string               `json:"title" form:"title"`
	Subtitle      *string               `json:"subtitle" form:"subtitle"`
	Content       *string               `json:"text" form:"text"`
}
