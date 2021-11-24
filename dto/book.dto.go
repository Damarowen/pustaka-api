package dto

import "time"

type BookRequest struct {
	Title       string ` json:"title" binding:"required" `
	Price       int    ` json:"price" binding:"required" `
	Description string ` json:"description" binding:"required" `
	Rating      int    ` json:"rating" binding:"required" `
	UserID      uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}

type BookResponse struct {
	Id          uint    ` json:"ID"`
	Title       string ` json:"TITLE" `
	Price       int    ` json:"PRICE"`
	Description string ` json:"DESC"`
	Rating      int    ` json:"RATING" `
}


type BookUpdateDTO struct {
	ID          uint `json:"id" form:"id" `
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Price       int    ` json:"price" binding:"required" `
	Rating      int    ` json:"rating" binding:"required" `
	UserID      uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
	CreatedAt time.Time 
}