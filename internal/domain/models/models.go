package models

type User struct { 
	Name   string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required, email"`
	Pass  int    `json:"pass" validate:"required"`
}

type Book struct {
	Id     string `json:"id"`
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
}